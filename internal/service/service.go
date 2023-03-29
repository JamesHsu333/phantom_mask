package service

import (
	"context"

	"github.com/JamesHsu333/kdan/internal/service/pharmacy"
	"github.com/JamesHsu333/kdan/internal/service/user"
	kdanProto "github.com/JamesHsu333/kdan/proto/kdan"
)

type kdanService struct {
	user.UserEndpoint
	pharmacy.PharmacyEndpoint
	kdanProto.UnimplementedKdanServiceServer
}

func (k *kdanService) GetPharmaciesByTime(ctx context.Context, r *kdanProto.GetPharmaciesByTimeRequest) (*kdanProto.GetPharmaciesByTimeResponse, error) {
	return k.PharmacyEndpoint.GetPharmaciesByTime(ctx, r)
}

func (k *kdanService) GetSoldMasksByPharmacy(ctx context.Context, r *kdanProto.GetSoldMasksByPharmacyRequest) (*kdanProto.GetSoldMasksByPharmacyResponse, error) {
	return k.PharmacyEndpoint.GetSoldMasksByPharmacy(ctx, r)
}

func (k *kdanService) GetPharmaciesMaskCountsByMaskPriceRange(ctx context.Context, r *kdanProto.GetPharmaciesMaskCountsByMaskPriceRangeRequest) (*kdanProto.GetPharmaciesMaskCountsByMaskPriceRangeResponse, error) {
	return k.PharmacyEndpoint.GetPharmaciesMaskCountsByMaskPriceRange(ctx, r)
}

func (k *kdanService) GetPharmaciesByNameRelevancy(ctx context.Context, r *kdanProto.GetPharmaciesByNameRelevancyRequest) (*kdanProto.GetPharmaciesByNameRelevancyResponse, error) {
	return k.PharmacyEndpoint.GetPharmaciesByNameRelevancy(ctx, r)
}

func (k *kdanService) GetMasksByNameRelevancy(ctx context.Context, r *kdanProto.GetMasksByNameRelevancyRequest) (*kdanProto.GetMasksByNameRelevancyResponse, error) {
	return k.PharmacyEndpoint.GetMasksByNameRelevancy(ctx, r)
}

func (k *kdanService) GetTopXUsersTransactionByDateRange(ctx context.Context, r *kdanProto.GetTopXUsersTransactionByDateRangeRequest) (*kdanProto.GetTopXUsersTransactionByDateRangeResponse, error) {
	return k.UserEndpoint.GetTopXUsersTransactionByDateRange(ctx, r)
}

func (k *kdanService) GetAggTransactionsByDateRange(ctx context.Context, r *kdanProto.GetAggTransactionsByDateRangeRequest) (*kdanProto.GetAggTransactionsByDateRangeResponse, error) {
	return k.UserEndpoint.GetAggTransactionsByDateRange(ctx, r)
}

func (k *kdanService) PurchaseMaskFromPharmacy(ctx context.Context, r *kdanProto.PurchaseMaskFromPharmacyRequest) (*kdanProto.PurchaseMaskFromPharmacyResponse, error) {
	return k.UserEndpoint.PurchaseMaskFromPharmacy(ctx, r)
}
