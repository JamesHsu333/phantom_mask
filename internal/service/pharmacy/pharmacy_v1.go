package pharmacy

import (
	"context"
	"database/sql"
	"errors"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/JamesHsu333/kdan/config"
	data "github.com/JamesHsu333/kdan/internal/data/test"
	"github.com/JamesHsu333/kdan/pkg/grpc_errors"
	"github.com/JamesHsu333/kdan/pkg/logger"
	kdanProto "github.com/JamesHsu333/kdan/proto/kdan"
	pharmacyProto "github.com/JamesHsu333/kdan/proto/pharmacy"
)

var sortedBy = map[kdanProto.GetSoldMasksByPharmacyRequest_SortedBy]string{
	kdanProto.GetSoldMasksByPharmacyRequest_mask_name:  "mask_name",
	kdanProto.GetSoldMasksByPharmacyRequest_mask_price: "mask_price",
}

func NewPharmacyEndpoint(logger logger.Logger, cfg *config.Config, pharmacyUC *pharmacyUC) *PharmacyEndpoint {
	return &PharmacyEndpoint{logger: logger, cfg: cfg, pharmacyUC: pharmacyUC}
}

func (p *PharmacyEndpoint) GetPharmaciesByTime(ctx context.Context, r *kdanProto.GetPharmaciesByTimeRequest) (*kdanProto.GetPharmaciesByTimeResponse, error) {
	totalMinutes, err := parseTimeToMinutes(r.GetDay(), r.GetHour(), r.GetMinute())
	if err != nil {
		p.logger.Errorf("pharmacyUC.parseTimeToMinutes: %v", err)
		return nil, err
	}

	pharmacies, err := p.pharmacyUC.GetPharmaciesByTime(ctx, totalMinutes)
	if err != nil {
		p.logger.Errorf("pharmacyUC.GetPharmaciesByTime: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "GetPharmaciesByTime: %v", err)
	}

	return &kdanProto.GetPharmaciesByTimeResponse{Pharmacies: p.pharmacyListModelToProto(pharmacies)}, err
}

func (p *PharmacyEndpoint) GetSoldMasksByPharmacy(ctx context.Context, r *kdanProto.GetSoldMasksByPharmacyRequest) (*kdanProto.GetSoldMasksByPharmacyResponse, error) {
	arg := data.GetSoldMasksByPharmacyParams{
		PharmacyName: r.GetName(),
		SortedBy:     sortedBy[r.GetSortedBy()],
		Reverse:      r.GetOrderBy() == kdanProto.Order_desc,
	}

	soldMasks, err := p.pharmacyUC.GetSoldMasksByPharmacy(ctx, arg)
	if err != nil {
		p.logger.Errorf("pharmacyUC.GetSoldMasksByPharmacy: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "GetSoldMasksByPharmacy: %v", err)
	}

	soldMasksProto := make([]*kdanProto.GetSoldMasksByPharmacyResponseSoldMask, 0, len(soldMasks))
	for _, s := range soldMasks {
		soldMaskProto := &kdanProto.GetSoldMasksByPharmacyResponseSoldMask{
			MaskId:       s.MaskID,
			MaskName:     s.Mask,
			PharmacyId:   s.PharmacyID,
			PharmacyName: s.PharmacyName,
			Price:        float32(s.TransactionAmount),
			SoldAt:       timestamppb.New(s.TransactionDate),
		}
		soldMasksProto = append(soldMasksProto, soldMaskProto)
	}

	if len(soldMasksProto) == 0 {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(sql.ErrNoRows), "pharmacyUC.GetSoldMasksByPharmacy : %v", sql.ErrNoRows)
	}

	return &kdanProto.GetSoldMasksByPharmacyResponse{SoldMasks: soldMasksProto}, err
}

func (p *PharmacyEndpoint) GetPharmaciesMaskCountsByMaskPriceRange(ctx context.Context, r *kdanProto.GetPharmaciesMaskCountsByMaskPriceRangeRequest) (*kdanProto.GetPharmaciesMaskCountsByMaskPriceRangeResponse, error) {
	startPrice := r.GetStartPrice()
	endPrice := r.GetEndPrice()

	if startPrice > endPrice {
		err := errors.New("Invalid price range format")
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "GetPharmaciesMaskCountsByMaskPriceRange: %v", err)
	}

	if r.GetMaskTypeCount() <= 0 {
		err := errors.New("Invalid mask type count")
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "GetPharmaciesMaskCountsByMaskPriceRange: %v", err)
	}

	arg := data.GetPharmaciesMaskCountsByMaskPriceRangeParams{
		StartPrice:    float64(startPrice),
		EndPrice:      float64(endPrice),
		MoreThan:      r.GetMoreThan(),
		MaskTypeCount: r.GetMaskTypeCount(),
	}

	maskCountPharms, err := p.pharmacyUC.GetPharmaciesMaskCountsByMaskPriceRange(ctx, arg)
	if err != nil {
		p.logger.Errorf("pharmacyUC.GetPharmaciesMaskCountsByMaskPriceRange: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "GetPharmaciesMaskCountsByMaskPriceRange: %v", err)
	}

	maskCountPharmsProto := make([]*kdanProto.GetPharmaciesMaskCountsByMaskPriceRangeResponsePharmaciesMaskCount, 0, len(maskCountPharms))
	for _, m := range maskCountPharms {
		maskCountPerPharm := &kdanProto.GetPharmaciesMaskCountsByMaskPriceRangeResponsePharmaciesMaskCount{
			PharmacyId:    m.PharmacyID,
			PharmacyName:  m.PharmacyName,
			MaskTypeCount: int32(m.MaskTypeCounts),
		}
		maskCountPharmsProto = append(maskCountPharmsProto, maskCountPerPharm)
	}

	if len(maskCountPharmsProto) == 0 {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(sql.ErrNoRows), "pharmacyUC.GetPharmaciesMaskCountsByMaskPriceRange : %v", sql.ErrNoRows)
	}

	return &kdanProto.GetPharmaciesMaskCountsByMaskPriceRangeResponse{PharmaciesMaskCounts: maskCountPharmsProto}, err
}

func (p *PharmacyEndpoint) GetPharmaciesByNameRelevancy(ctx context.Context, r *kdanProto.GetPharmaciesByNameRelevancyRequest) (*kdanProto.GetPharmaciesByNameRelevancyResponse, error) {
	if r.GetName() == "" {
		err := errors.New("invalid arguments: must provide a name")
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "pharmacyUC.GetPharmaciesByNameRelevancy: %v", err)
	}

	pharmacies, err := p.pharmacyUC.GetPharmaciesByNameRelevancy(ctx, r.GetName())
	if err != nil {
		p.logger.Errorf("pharmacyUC.GetPharmaciesByNameRelevancy: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "pharmacyUC.GetPharmaciesByNameRelevancy: %v", err)
	}

	if len(pharmacies) == 0 {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(sql.ErrNoRows), "pharmacyUC.GetPharmaciesByNameRelevancy : %v", sql.ErrNoRows)
	}

	return &kdanProto.GetPharmaciesByNameRelevancyResponse{Pharmacies: p.pharmacyListModelToProto(pharmacies)}, nil
}

func (p *PharmacyEndpoint) GetMasksByNameRelevancy(ctx context.Context, r *kdanProto.GetMasksByNameRelevancyRequest) (*kdanProto.GetMasksByNameRelevancyResponse, error) {
	if r.GetName() == "" {
		err := errors.New("invalid arguments: must provide a name")
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "pharmacyUC.GetMasksByNameRelevancy: %v", err)
	}

	masks, err := p.pharmacyUC.GetMasksByNameRelevancy(ctx, r.GetName())
	if err != nil {
		p.logger.Errorf("pharmacyUC.GetMasksByNameRelevancy : %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "pharmacyUC.GetMasksByNameRelevancy : %v", err)
	}

	masksProto := make([]*pharmacyProto.Mask, 0, len(masks))
	for _, mask := range masks {
		maskProto := &pharmacyProto.Mask{
			Id:   mask.ID,
			Name: mask.Name,
		}
		masksProto = append(masksProto, maskProto)
	}

	if len(masksProto) == 0 {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(sql.ErrNoRows), "pharmacyUC.GetMasksByNameRelevancy : %v", sql.ErrNoRows)
	}

	return &kdanProto.GetMasksByNameRelevancyResponse{Masks: masksProto}, nil
}

func (p *PharmacyEndpoint) pharmacyModelToProto(pharmacy *data.Pharmacy) *pharmacyProto.Pharmacy {
	pharmacyProto := &pharmacyProto.Pharmacy{
		Id:           pharmacy.ID,
		Name:         pharmacy.Name,
		OpeningHours: pharmacy.OpeningHoursDescription,
		CashBalance:  float32(*pharmacy.CashBalance),
	}

	return pharmacyProto
}

func (p *PharmacyEndpoint) pharmacyListModelToProto(pharmacies []data.Pharmacy) []*pharmacyProto.Pharmacy {
	pharmacyList := make([]*pharmacyProto.Pharmacy, 0, len(pharmacies))

	for _, pharmacy := range pharmacies {
		pharmacyProto := p.pharmacyModelToProto(&pharmacy)
		pharmacyList = append(pharmacyList, pharmacyProto)
	}

	return pharmacyList
}

func parseTimeToMinutes(day kdanProto.DayOfWeek, hour int32, minute int32) (int32, error) {
	if day == kdanProto.DayOfWeek_unspecified {
		return 0, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(grpc_errors.ErrInvalidDayTimeFormat), "day unspecified: %v", grpc_errors.ErrInvalidDayTimeFormat)
	}

	if hour > 24 || hour < 0 {
		return 0, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(grpc_errors.ErrInvalidDayTimeFormat), "invalid hour format: %v", grpc_errors.ErrInvalidDayTimeFormat)
	}

	if minute > 60 || minute < 0 {
		return 0, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(grpc_errors.ErrInvalidDayTimeFormat), "invalid minute format: %v", grpc_errors.ErrInvalidDayTimeFormat)
	}

	weekDay := int32(day)
	totalMinutes := (weekDay-1)*24*60 + hour*60 + minute

	return totalMinutes, nil
}
