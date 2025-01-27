// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: pharmacy.sql

package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMask = `-- name: CreateMask :one
INSERT INTO masks (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateMask(ctx context.Context, name string) (Mask, error) {
	row := q.db.QueryRow(ctx, createMask, name)
	var i Mask
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createMaskPrice = `-- name: CreateMaskPrice :exec
INSERT INTO mask_prices (pharmacy_id, mask_id, price)
VALUES ($1, $2, $3)
RETURNING mask_id, pharmacy_id, price
`

type CreateMaskPriceParams struct {
	PharmacyID int32   `json:"pharmacy_id"`
	MaskID     int32   `json:"mask_id"`
	Price      float64 `json:"price"`
}

func (q *Queries) CreateMaskPrice(ctx context.Context, arg CreateMaskPriceParams) error {
	_, err := q.db.Exec(ctx, createMaskPrice, arg.PharmacyID, arg.MaskID, arg.Price)
	return err
}

const createPharmacy = `-- name: CreatePharmacy :one
INSERT INTO pharmacies (name, opening_hours, opening_hours_description, cash_balance)
VALUES ($1, $2, $3, $4)
RETURNING id, name, opening_hours, opening_hours_description, cash_balance
`

type CreatePharmacyParams struct {
	Name                    string                                       `json:"name"`
	OpeningHours            pgtype.Multirange[pgtype.Range[pgtype.Int4]] `json:"opening_hours"`
	OpeningHoursDescription string                                       `json:"opening_hours_description"`
	CashBalance             *float64                                     `json:"cash_balance"`
}

func (q *Queries) CreatePharmacy(ctx context.Context, arg CreatePharmacyParams) (Pharmacy, error) {
	row := q.db.QueryRow(ctx, createPharmacy,
		arg.Name,
		arg.OpeningHours,
		arg.OpeningHoursDescription,
		arg.CashBalance,
	)
	var i Pharmacy
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OpeningHours,
		&i.OpeningHoursDescription,
		&i.CashBalance,
	)
	return i, err
}

const getMaskPriceById = `-- name: GetMaskPriceById :one
SELECT price FROM mask_prices
WHERE $1::int = mask_id AND $2::int = pharmacy_id
`

type GetMaskPriceByIdParams struct {
	MaskID     int32 `json:"mask_id"`
	PharmacyID int32 `json:"pharmacy_id"`
}

func (q *Queries) GetMaskPriceById(ctx context.Context, arg GetMaskPriceByIdParams) (float64, error) {
	row := q.db.QueryRow(ctx, getMaskPriceById, arg.MaskID, arg.PharmacyID)
	var price float64
	err := row.Scan(&price)
	return price, err
}

const getMasksByNameRelevancy = `-- name: GetMasksByNameRelevancy :many
SELECT id, name FROM masks
WHERE SIMILARITY(name,$1::text) > 0.2
ORDER BY SIMILARITY(name,$1::text) DESC
`

func (q *Queries) GetMasksByNameRelevancy(ctx context.Context, name string) ([]Mask, error) {
	rows, err := q.db.Query(ctx, getMasksByNameRelevancy, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Mask
	for rows.Next() {
		var i Mask
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPharmaciesByNameRelevancy = `-- name: GetPharmaciesByNameRelevancy :many
SELECT id, name, opening_hours, opening_hours_description, cash_balance FROM pharmacies
WHERE SIMILARITY(name,$1::text) > 0.2
ORDER BY SIMILARITY(name,$1::text) DESC
`

func (q *Queries) GetPharmaciesByNameRelevancy(ctx context.Context, name string) ([]Pharmacy, error) {
	rows, err := q.db.Query(ctx, getPharmaciesByNameRelevancy, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pharmacy
	for rows.Next() {
		var i Pharmacy
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.OpeningHours,
			&i.OpeningHoursDescription,
			&i.CashBalance,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPharmaciesByTime = `-- name: GetPharmaciesByTime :many
SELECT id, name, opening_hours, opening_hours_description, cash_balance FROM pharmacies
WHERE $1::int <@ opening_hours
`

func (q *Queries) GetPharmaciesByTime(ctx context.Context, time int32) ([]Pharmacy, error) {
	rows, err := q.db.Query(ctx, getPharmaciesByTime, time)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pharmacy
	for rows.Next() {
		var i Pharmacy
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.OpeningHours,
			&i.OpeningHoursDescription,
			&i.CashBalance,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPharmaciesMaskCountsByMaskPriceRange = `-- name: GetPharmaciesMaskCountsByMaskPriceRange :many
SELECT p.id AS pharmacy_id, p.name AS pharmacy_name, COUNT(distinct m.name) AS mask_type_counts FROM mask_prices mp
INNER JOIN masks m ON mp.mask_id = m.id
INNER JOIN pharmacies p ON mp.pharmacy_id = p.id
WHERE mp.price BETWEEN $1::NUMERIC AND $2::NUMERIC
GROUP BY p.id, p.name
HAVING CASE
    WHEN $3::boolean THEN $4::int <= COUNT(distinct m.name)
    WHEN NOT $3::boolean THEN $4::int >= COUNT(distinct m.name)
END
`

type GetPharmaciesMaskCountsByMaskPriceRangeParams struct {
	StartPrice    float64 `json:"start_price"`
	EndPrice      float64 `json:"end_price"`
	MoreThan      bool    `json:"more_than"`
	MaskTypeCount int32   `json:"mask_type_count"`
}

type GetPharmaciesMaskCountsByMaskPriceRangeRow struct {
	PharmacyID     int32  `json:"pharmacy_id"`
	PharmacyName   string `json:"pharmacy_name"`
	MaskTypeCounts int64  `json:"mask_type_counts"`
}

func (q *Queries) GetPharmaciesMaskCountsByMaskPriceRange(ctx context.Context, arg GetPharmaciesMaskCountsByMaskPriceRangeParams) ([]GetPharmaciesMaskCountsByMaskPriceRangeRow, error) {
	rows, err := q.db.Query(ctx, getPharmaciesMaskCountsByMaskPriceRange,
		arg.StartPrice,
		arg.EndPrice,
		arg.MoreThan,
		arg.MaskTypeCount,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPharmaciesMaskCountsByMaskPriceRangeRow
	for rows.Next() {
		var i GetPharmaciesMaskCountsByMaskPriceRangeRow
		if err := rows.Scan(&i.PharmacyID, &i.PharmacyName, &i.MaskTypeCounts); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSoldMasksByPharmacy = `-- name: GetSoldMasksByPharmacy :many
SELECT ph.transaction_amount, ph.transaction_date, m.name AS mask, m.id AS mask_id, p.name AS pharmacy_name, p.id AS pharmacy_id
FROM purchase_histories ph
INNER JOIN masks m ON ph.mask_id = m.id
INNER JOIN pharmacies p ON ph.pharmacy_id = p.id
WHERE $1::text = p.name
ORDER BY CASE
    WHEN NOT $2::boolean AND $3::text = 'mask_name' THEN m.name
    WHEN NOT $2::boolean AND $3::text = 'mask_price' THEN ph.transaction_amount::varchar
END ASC, CASE
    WHEN $2::boolean AND $3::text = 'mask_name' THEN m.name
    WHEN $2::boolean AND $3::text = 'mask_price' THEN ph.transaction_amount::varchar
END  DESC
`

type GetSoldMasksByPharmacyParams struct {
	PharmacyName string `json:"pharmacy_name"`
	Reverse      bool   `json:"reverse"`
	SortedBy     string `json:"sorted_by"`
}

type GetSoldMasksByPharmacyRow struct {
	TransactionAmount float64   `json:"transaction_amount"`
	TransactionDate   time.Time `json:"transaction_date"`
	Mask              string    `json:"mask"`
	MaskID            int32     `json:"mask_id"`
	PharmacyName      string    `json:"pharmacy_name"`
	PharmacyID        int32     `json:"pharmacy_id"`
}

func (q *Queries) GetSoldMasksByPharmacy(ctx context.Context, arg GetSoldMasksByPharmacyParams) ([]GetSoldMasksByPharmacyRow, error) {
	rows, err := q.db.Query(ctx, getSoldMasksByPharmacy, arg.PharmacyName, arg.Reverse, arg.SortedBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSoldMasksByPharmacyRow
	for rows.Next() {
		var i GetSoldMasksByPharmacyRow
		if err := rows.Scan(
			&i.TransactionAmount,
			&i.TransactionDate,
			&i.Mask,
			&i.MaskID,
			&i.PharmacyName,
			&i.PharmacyID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
