-- name: CreatePharmacy :one
INSERT INTO pharmacies (name, opening_hours, opening_hours_description, cash_balance)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateMask :one
INSERT INTO masks (name)
VALUES ($1)
RETURNING *;

-- name: CreateMaskPrice :exec
INSERT INTO mask_prices (pharmacy_id, mask_id, price)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPharmaciesByTime :many
SELECT * FROM pharmacies
WHERE @time::int <@ opening_hours;

-- name: GetSoldMasksByPharmacy :many
SELECT ph.transaction_amount, ph.transaction_date, m.name AS mask, m.id AS mask_id, p.name AS pharmacy_name, p.id AS pharmacy_id
FROM purchase_histories ph
INNER JOIN masks m ON ph.mask_id = m.id
INNER JOIN pharmacies p ON ph.pharmacy_id = p.id
WHERE @pharmacy_name::text = p.name
ORDER BY CASE
    WHEN NOT @reverse::boolean AND @sorted_by::text = 'mask_name' THEN m.name
    WHEN NOT @reverse::boolean AND @sorted_by::text = 'mask_price' THEN ph.transaction_amount::varchar
END ASC, CASE
    WHEN @reverse::boolean AND @sorted_by::text = 'mask_name' THEN m.name
    WHEN @reverse::boolean AND @sorted_by::text = 'mask_price' THEN ph.transaction_amount::varchar
END  DESC;

-- name: GetPharmaciesMaskCountsByMaskPriceRange :many
SELECT p.id AS pharmacy_id, p.name AS pharmacy_name, COUNT(distinct m.name) AS mask_type_counts FROM mask_prices mp
INNER JOIN masks m ON mp.mask_id = m.id
INNER JOIN pharmacies p ON mp.pharmacy_id = p.id
WHERE mp.price BETWEEN @start_price::NUMERIC AND @end_price::NUMERIC
GROUP BY p.id, p.name
HAVING CASE
    WHEN @more_than::boolean THEN @mask_type_count::int <= COUNT(distinct m.name)
    WHEN NOT @more_than::boolean THEN @mask_type_count::int >= COUNT(distinct m.name)
END;

-- name: GetPharmaciesByNameRelevancy :many
SELECT * FROM pharmacies
WHERE SIMILARITY(name,@name::text) > 0.2
ORDER BY SIMILARITY(name,@name::text) DESC;

-- name: GetMasksByNameRelevancy :many
SELECT * FROM masks
WHERE SIMILARITY(name,@name::text) > 0.2
ORDER BY SIMILARITY(name,@name::text) DESC;

-- name: GetMaskPriceById :one
SELECT price FROM mask_prices
WHERE @mask_id::int = mask_id AND @pharmacy_id::int = pharmacy_id;

