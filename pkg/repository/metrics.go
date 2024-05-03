package repository

import (
	"context"

	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
	"github.com/rahul7668gupta/go-url-shortner/pkg/dto"
)

func (r *ShortnerRepo) GetMetrics(ctx context.Context) ([]dto.Metrics, error) {
	domainCounts, err := r.rdb.ZRevRangeWithScores(ctx, constants.DOMAIN_COUNTER, 0, 2).Result()
	if err != nil {
		r.logger.Sugar().Errorf("error retrieving metrics %s", err.Error())
		return nil, err
	}

	var topDomains []dto.Metrics
	for _, domainCount := range domainCounts {
		topDomains = append(topDomains, dto.Metrics{
			DomainName: domainCount.Member.(string),
			Count:      int64(domainCount.Score),
		})
	}
	return topDomains, nil
}
