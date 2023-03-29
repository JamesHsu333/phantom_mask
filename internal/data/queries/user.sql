-- name: CreateUser :one
INSERT INTO users (name, cash_balance)
VALUES ($1, $2)
RETURNING *;

-- name: CreatePurchase :one
INSERT INTO purchase_histories (user_id, pharmacy_id, mask_id, transaction_amount, transaction_date)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTopXUsersTransactionByDateRange :many
SELECT u.id AS user_id, u.name AS user_name, SUM(ph.transaction_amount)::NUMERIC AS total_transaction_amount FROM purchase_histories ph
INNER JOIN users u ON ph.user_id = u.id
WHERE ph.transaction_date BETWEEN @start_date::timestamptz AND @end_date::timestamptz
GROUP BY u.id, u.name
ORDER BY total_transaction_amount desc
LIMIT @top::int;

-- name: PurchaseMaskFromPharmacy :one
INSERT INTO purchase_histories (user_id, pharmacy_id, mask_id, transaction_amount, transaction_date)
VALUES (@user_id::int, @pharmacy_id::int, @mask_id::int, @transaction_amount::NUMERIC, now()) RETURNING *;

-- name: UpdateUserCashBalance :exec
UPDATE users SET cash_balance = cash_balance - @transaction_amount::NUMERIC WHERE id = @user_id::int;

-- name: UpdatePharmacyCashBalance :exec
UPDATE pharmacies SET cash_balance = cash_balance + @transaction_amount::NUMERIC WHERE id = @pharmacy_id::int;

-- name: GetAggTransactionsByDateRange :many
SELECT m.id AS mask_id, m.name AS mask_name, COUNT(m.name) AS sold_mask_count, SUM(ph.transaction_amount)::NUMERIC AS total_transaction_amount FROM purchase_histories ph
INNER JOIN masks m ON ph.mask_id = m.id
WHERE ph.transaction_date BETWEEN @start_date::timestamptz AND @end_date::timestamptz
GROUP BY m.id, m.name;