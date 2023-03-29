package user

import (
	"context"

	data "github.com/JamesHsu333/kdan/internal/data/test"
	"github.com/JamesHsu333/kdan/pkg/logger"
	"github.com/jackc/pgx/v5"
)

// Auth UseCase
type userUC struct {
	querier *data.Queries
	pgconn  *pgx.Conn
	logger  logger.Logger
}

// Auth UseCase constructor
func NewUserUC(querier *data.Queries, pgconn *pgx.Conn, log logger.Logger) *userUC {
	return &userUC{querier: querier, pgconn: pgconn, logger: log}
}

func (u *userUC) GetTopXUsersTransactionByDateRange(ctx context.Context, arg data.GetTopXUsersTransactionByDateRangeParams) ([]data.GetTopXUsersTransactionByDateRangeRow, error) {
	results, err := u.querier.GetTopXUsersTransactionByDateRange(ctx, arg)
	if err != nil {
		return nil, err
	}

	return results, err
}

func (u *userUC) PurchaseMaskFromPharmacy(ctx context.Context, userId int32, pharmacyId int32, maskId int32) (*data.PurchaseHistory, error) {
	maskPrice, err := u.querier.GetMaskPriceById(ctx, data.GetMaskPriceByIdParams{
		MaskID:     maskId,
		PharmacyID: pharmacyId,
	})
	if err != nil || maskPrice == 0 {
		return nil, err
	}

	tx, err := u.pgconn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	qtx := u.querier.WithTx(tx)

	purchaseHistory, err := qtx.PurchaseMaskFromPharmacy(ctx, data.PurchaseMaskFromPharmacyParams{
		UserID:            userId,
		PharmacyID:        pharmacyId,
		MaskID:            maskId,
		TransactionAmount: maskPrice,
	})
	if err != nil {
		return nil, err
	}

	err = qtx.UpdateUserCashBalance(ctx, data.UpdateUserCashBalanceParams{
		UserID:            userId,
		TransactionAmount: maskPrice,
	})
	if err != nil {
		return nil, err
	}

	err = qtx.UpdatePharmacyCashBalance(ctx, data.UpdatePharmacyCashBalanceParams{
		PharmacyID:        pharmacyId,
		TransactionAmount: maskPrice,
	})
	if err != nil {
		return nil, err
	}

	return &purchaseHistory, err
}

func (u *userUC) GetAggTransactionsByDateRange(ctx context.Context, arg data.GetAggTransactionsByDateRangeParams) ([]data.GetAggTransactionsByDateRangeRow, error) {
	results, err := u.querier.GetAggTransactionsByDateRange(ctx, arg)
	if err != nil {
		return nil, err
	}

	return results, err
}
