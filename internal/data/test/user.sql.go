// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user.sql

package data

import (
	"context"
	"time"
)

const createPurchase = `-- name: CreatePurchase :one
INSERT INTO purchase_histories (user_id, pharmacy_id, mask_id, transaction_amount, transaction_date)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, pharmacy_id, mask_id, transaction_amount, transaction_date
`

type CreatePurchaseParams struct {
	UserID            int32     `json:"user_id"`
	PharmacyID        int32     `json:"pharmacy_id"`
	MaskID            int32     `json:"mask_id"`
	TransactionAmount float64   `json:"transaction_amount"`
	TransactionDate   time.Time `json:"transaction_date"`
}

func (q *Queries) CreatePurchase(ctx context.Context, arg CreatePurchaseParams) (PurchaseHistory, error) {
	row := q.db.QueryRow(ctx, createPurchase,
		arg.UserID,
		arg.PharmacyID,
		arg.MaskID,
		arg.TransactionAmount,
		arg.TransactionDate,
	)
	var i PurchaseHistory
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PharmacyID,
		&i.MaskID,
		&i.TransactionAmount,
		&i.TransactionDate,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, cash_balance)
VALUES ($1, $2)
RETURNING id, name, cash_balance
`

type CreateUserParams struct {
	Name        string   `json:"name"`
	CashBalance *float64 `json:"cash_balance"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Name, arg.CashBalance)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.CashBalance)
	return i, err
}

const getAggTransactionsByDateRange = `-- name: GetAggTransactionsByDateRange :many
SELECT m.id AS mask_id, m.name AS mask_name, COUNT(m.name) AS sold_mask_count, SUM(ph.transaction_amount)::NUMERIC AS total_transaction_amount FROM purchase_histories ph
INNER JOIN masks m ON ph.mask_id = m.id
WHERE ph.transaction_date BETWEEN $1::timestamptz AND $2::timestamptz
GROUP BY m.id, m.name
`

type GetAggTransactionsByDateRangeParams struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type GetAggTransactionsByDateRangeRow struct {
	MaskID                 int32   `json:"mask_id"`
	MaskName               string  `json:"mask_name"`
	SoldMaskCount          int64   `json:"sold_mask_count"`
	TotalTransactionAmount float64 `json:"total_transaction_amount"`
}

func (q *Queries) GetAggTransactionsByDateRange(ctx context.Context, arg GetAggTransactionsByDateRangeParams) ([]GetAggTransactionsByDateRangeRow, error) {
	rows, err := q.db.Query(ctx, getAggTransactionsByDateRange, arg.StartDate, arg.EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAggTransactionsByDateRangeRow
	for rows.Next() {
		var i GetAggTransactionsByDateRangeRow
		if err := rows.Scan(
			&i.MaskID,
			&i.MaskName,
			&i.SoldMaskCount,
			&i.TotalTransactionAmount,
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

const getTopXUsersTransactionByDateRange = `-- name: GetTopXUsersTransactionByDateRange :many
SELECT u.id AS user_id, u.name AS user_name, SUM(ph.transaction_amount)::NUMERIC AS total_transaction_amount FROM purchase_histories ph
INNER JOIN users u ON ph.user_id = u.id
WHERE ph.transaction_date BETWEEN $1::timestamptz AND $2::timestamptz
GROUP BY u.id, u.name
ORDER BY total_transaction_amount desc
LIMIT $3::int
`

type GetTopXUsersTransactionByDateRangeParams struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Top       int32     `json:"top"`
}

type GetTopXUsersTransactionByDateRangeRow struct {
	UserID                 int32   `json:"user_id"`
	UserName               string  `json:"user_name"`
	TotalTransactionAmount float64 `json:"total_transaction_amount"`
}

func (q *Queries) GetTopXUsersTransactionByDateRange(ctx context.Context, arg GetTopXUsersTransactionByDateRangeParams) ([]GetTopXUsersTransactionByDateRangeRow, error) {
	rows, err := q.db.Query(ctx, getTopXUsersTransactionByDateRange, arg.StartDate, arg.EndDate, arg.Top)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTopXUsersTransactionByDateRangeRow
	for rows.Next() {
		var i GetTopXUsersTransactionByDateRangeRow
		if err := rows.Scan(&i.UserID, &i.UserName, &i.TotalTransactionAmount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const purchaseMaskFromPharmacy = `-- name: PurchaseMaskFromPharmacy :one
INSERT INTO purchase_histories (user_id, pharmacy_id, mask_id, transaction_amount, transaction_date)
VALUES ($1::int, $2::int, $3::int, $4::NUMERIC, now()) RETURNING id, user_id, pharmacy_id, mask_id, transaction_amount, transaction_date
`

type PurchaseMaskFromPharmacyParams struct {
	UserID            int32   `json:"user_id"`
	PharmacyID        int32   `json:"pharmacy_id"`
	MaskID            int32   `json:"mask_id"`
	TransactionAmount float64 `json:"transaction_amount"`
}

func (q *Queries) PurchaseMaskFromPharmacy(ctx context.Context, arg PurchaseMaskFromPharmacyParams) (PurchaseHistory, error) {
	row := q.db.QueryRow(ctx, purchaseMaskFromPharmacy,
		arg.UserID,
		arg.PharmacyID,
		arg.MaskID,
		arg.TransactionAmount,
	)
	var i PurchaseHistory
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PharmacyID,
		&i.MaskID,
		&i.TransactionAmount,
		&i.TransactionDate,
	)
	return i, err
}

const updatePharmacyCashBalance = `-- name: UpdatePharmacyCashBalance :exec
UPDATE pharmacies SET cash_balance = cash_balance + $1::NUMERIC WHERE id = $2::int
`

type UpdatePharmacyCashBalanceParams struct {
	TransactionAmount float64 `json:"transaction_amount"`
	PharmacyID        int32   `json:"pharmacy_id"`
}

func (q *Queries) UpdatePharmacyCashBalance(ctx context.Context, arg UpdatePharmacyCashBalanceParams) error {
	_, err := q.db.Exec(ctx, updatePharmacyCashBalance, arg.TransactionAmount, arg.PharmacyID)
	return err
}

const updateUserCashBalance = `-- name: UpdateUserCashBalance :exec
UPDATE users SET cash_balance = cash_balance - $1::NUMERIC WHERE id = $2::int
`

type UpdateUserCashBalanceParams struct {
	TransactionAmount float64 `json:"transaction_amount"`
	UserID            int32   `json:"user_id"`
}

func (q *Queries) UpdateUserCashBalance(ctx context.Context, arg UpdateUserCashBalanceParams) error {
	_, err := q.db.Exec(ctx, updateUserCashBalance, arg.TransactionAmount, arg.UserID)
	return err
}
