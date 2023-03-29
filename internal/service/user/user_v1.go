package user

import (
	"context"
	"errors"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/JamesHsu333/kdan/config"
	data "github.com/JamesHsu333/kdan/internal/data/test"
	"github.com/JamesHsu333/kdan/pkg/grpc_errors"
	"github.com/JamesHsu333/kdan/pkg/logger"
	kdanProto "github.com/JamesHsu333/kdan/proto/kdan"
)

func NewUserEndpoint(logger logger.Logger, cfg *config.Config, userUC *userUC) *UserEndpoint {
	return &UserEndpoint{logger: logger, cfg: cfg, userUC: userUC}
}

func (u *UserEndpoint) GetTopXUsersTransactionByDateRange(ctx context.Context, r *kdanProto.GetTopXUsersTransactionByDateRangeRequest) (*kdanProto.GetTopXUsersTransactionByDateRangeResponse, error) {
	topCount := r.GetSize()
	if topCount == 0 {
		topCount = 10
	}

	startDate := r.GetStartAt()
	endDate := r.GetEndAt()
	if startDate == nil || endDate == nil || startDate == endDate {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(grpc_errors.ErrInvalidDayTimeFormat), "GetTopXUsersTransactionByDateRange: %v", grpc_errors.ErrInvalidDayTimeFormat)
	}

	if r.GetStartAt().AsTime().After(r.GetEndAt().AsTime()) {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(grpc_errors.ErrInvalidDayTimeFormat), "GetTopXUsersTransactionByDateRange: %v", errors.New("Invalid Request: StartAt is after EndAt"))
	}

	arg := data.GetTopXUsersTransactionByDateRangeParams{
		StartDate: r.GetStartAt().AsTime(),
		EndDate:   r.GetEndAt().AsTime(),
		Top:       topCount,
	}

	u.logger.Infof("%+v", arg)

	users, err := u.userUC.GetTopXUsersTransactionByDateRange(ctx, arg)
	if err != nil {
		u.logger.Errorf("userUC.GetTopXUsersTransactionByDateRange: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "GetTopXUsersTransactionByDateRange: %v", err)
	}

	userTransProto := make([]*kdanProto.GetTopXUsersTransactionByDateRangeResponseUserTransaction, 0, len(users))
	for _, user := range users {
		userTranProto := &kdanProto.GetTopXUsersTransactionByDateRangeResponseUserTransaction{
			UserId:                 user.UserID,
			UserName:               user.UserName,
			TotalTransactionAmount: float32(user.TotalTransactionAmount),
		}
		userTransProto = append(userTransProto, userTranProto)
	}

	return &kdanProto.GetTopXUsersTransactionByDateRangeResponse{UserTransactions: userTransProto}, err
}

func (u *UserEndpoint) GetAggTransactionsByDateRange(ctx context.Context, r *kdanProto.GetAggTransactionsByDateRangeRequest) (*kdanProto.GetAggTransactionsByDateRangeResponse, error) {
	startDate := r.GetStartAt()
	endDate := r.GetEndAt()
	if startDate == nil || endDate == nil || startDate == endDate {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(grpc_errors.ErrInvalidDayTimeFormat), "GetTopXUsersTransactionByDateRange: %v", grpc_errors.ErrInvalidDayTimeFormat)
	}

	if r.GetStartAt().AsTime().After(r.GetEndAt().AsTime()) {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(grpc_errors.ErrInvalidDayTimeFormat), "GetTopXUsersTransactionByDateRange: %v", errors.New("Invalid Request: StartAt is after EndAt"))
	}

	arg := data.GetAggTransactionsByDateRangeParams{
		StartDate: r.GetStartAt().AsTime(),
		EndDate:   r.GetEndAt().AsTime(),
	}

	maskTrans, err := u.userUC.GetAggTransactionsByDateRange(ctx, arg)
	if err != nil {
		u.logger.Errorf("userUC.GetAggTransactionsByDateRange: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "GetAggTransactionsByDateRange: %v", err)
	}

	maskTransProto := make([]*kdanProto.GetAggTransactionsByDateRangeResponseMaskTransaction, 0, len(maskTrans))
	for _, mask := range maskTrans {
		maskTranProto := &kdanProto.GetAggTransactionsByDateRangeResponseMaskTransaction{
			MaskId:                 mask.MaskID,
			MaskName:               mask.MaskName,
			SoldMaskCount:          int32(mask.SoldMaskCount),
			TotalTransactionAmount: float32(mask.TotalTransactionAmount),
		}
		maskTransProto = append(maskTransProto, maskTranProto)
	}

	return &kdanProto.GetAggTransactionsByDateRangeResponse{MaskTransactions: maskTransProto}, err
}

func (u *UserEndpoint) PurchaseMaskFromPharmacy(ctx context.Context, r *kdanProto.PurchaseMaskFromPharmacyRequest) (*kdanProto.PurchaseMaskFromPharmacyResponse, error) {
	if r.GetUserId() == 0 || r.GetPharmacyId() == 0 || r.GetMaskId() == 0 {
		err := errors.New("invalid request")
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "PurchaseMaskFromPharmacy: %v", err)
	}

	purchaseHistory, err := u.userUC.PurchaseMaskFromPharmacy(ctx, r.GetUserId(), r.GetPharmacyId(), r.GetMaskId())
	if err != nil {
		u.logger.Errorf("userUC.PurchaseMaskFromPharmacy: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "PurchaseMaskFromPharmacy: %v", err)
	}

	return &kdanProto.PurchaseMaskFromPharmacyResponse{
		Id:                purchaseHistory.ID,
		UserId:            purchaseHistory.UserID,
		PharmacyId:        purchaseHistory.PharmacyID,
		MaskId:            purchaseHistory.MaskID,
		TransactionAmount: float32(purchaseHistory.TransactionAmount),
		TransactionDate:   timestamppb.New(purchaseHistory.TransactionDate),
	}, err
}
