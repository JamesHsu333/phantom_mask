package pharmacy

import (
	"context"

	data "github.com/JamesHsu333/kdan/internal/data/test"
	"github.com/JamesHsu333/kdan/pkg/logger"
)

type pharmacyUC struct {
	querier data.Querier
	logger  logger.Logger
}

func NewPharmacyUC(querier data.Querier, log logger.Logger) *pharmacyUC {
	return &pharmacyUC{querier: querier, logger: log}
}

func (p *pharmacyUC) GetPharmaciesByTime(ctx context.Context, time int32) ([]data.Pharmacy, error) {
	pharmacies, err := p.querier.GetPharmaciesByTime(ctx, time)
	if err != nil {
		p.logger.Errorf("pharmacyUC.querier.GetPharmaciesByTime: %v", err)
		return nil, err
	}

	return pharmacies, err
}

func (p *pharmacyUC) GetSoldMasksByPharmacy(ctx context.Context, arg data.GetSoldMasksByPharmacyParams) ([]data.GetSoldMasksByPharmacyRow, error) {
	results, err := p.querier.GetSoldMasksByPharmacy(ctx, arg)
	if err != nil {
		p.logger.Errorf("pharmacyUC.querier.GetSoldMasksByPharmacy: %v", err)
		return nil, err
	}

	return results, err
}

func (p *pharmacyUC) GetPharmaciesMaskCountsByMaskPriceRange(ctx context.Context, arg data.GetPharmaciesMaskCountsByMaskPriceRangeParams) ([]data.GetPharmaciesMaskCountsByMaskPriceRangeRow, error) {
	results, err := p.querier.GetPharmaciesMaskCountsByMaskPriceRange(ctx, arg)
	if err != nil {
		p.logger.Errorf("pharmacyUC.querier.GetPharmaciesMaskCountsByMaskPriceRange: %v", err)
		return nil, err
	}

	return results, err
}

func (p *pharmacyUC) GetPharmaciesByNameRelevancy(ctx context.Context, name string) ([]data.Pharmacy, error) {
	pharmacies, err := p.querier.GetPharmaciesByNameRelevancy(ctx, name)
	if err != nil {
		p.logger.Errorf("pharmacyUC.querier.GetPharmaiesByNameRelevancy: %v", err)
		return nil, err
	}

	return pharmacies, err
}

func (p *pharmacyUC) GetMasksByNameRelevancy(ctx context.Context, name string) ([]data.Mask, error) {
	masks, err := p.querier.GetMasksByNameRelevancy(ctx, name)
	if err != nil {
		p.logger.Errorf("pharmacyUC.querier.GetMasksByNameRelevancy: %v", err)
		return nil, err
	}

	return masks, err
}
