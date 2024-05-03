package repository

import (
	"context"
	"log"

	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
	"github.com/rahul7668gupta/go-url-shortner/pkg/dto"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ShortnerRepo struct {
	rdb    *redis.Client
	logger *zap.Logger
}

type IShortnerRepo interface {
	CreateShortCodeRecord(ctx context.Context, shortCode string, requestUrl string) error
	IncrementDomainCounter(ctx context.Context, domain string) error
	GetOriginalUrlForShortCode(ctx context.Context, shortCode string) (string, error)
	GetMetrics(ctx context.Context) ([]dto.Metrics, error)
	CreateIndexOnOriginalUrl(ctx context.Context, url string, shortCode string) error
	LookupURL(ctx context.Context, url string) (string, bool)
}

func NewShortnerRepository(rdb *redis.Client, logger *zap.Logger) *ShortnerRepo {
	return &ShortnerRepo{
		rdb:    rdb,
		logger: logger,
	}
}

func (r *ShortnerRepo) CreateShortCodeRecord(ctx context.Context, shortCode string, requestUrl string) error {
	err := r.rdb.Set(ctx, shortCode, requestUrl, 0).Err()
	return err
}

func (r *ShortnerRepo) IncrementDomainCounter(ctx context.Context, domain string) error {
	err := r.rdb.ZIncrBy(ctx, constants.DOMAIN_COUNTER, 1, domain).Err()
	return err
}

func (r *ShortnerRepo) CreateIndexOnOriginalUrl(ctx context.Context, url string, shortCode string) error {
	err := r.rdb.Set(ctx, "index:url:"+url, shortCode, 0).Err()
	return err
}

func (r *ShortnerRepo) LookupURL(ctx context.Context, url string) (string, bool) {
	// Attempt to find the url in the index
	shortenedPath, err := r.rdb.Get(ctx, "index:url:"+url).Result()
	if err == redis.Nil {
		// url not found, return false
		return "", false
	} else if err != nil {
		// Handle the error
		log.Printf("Error looking up url: %v", err)
		return "", false
	}
	// url found, return the shortened path and true
	return shortenedPath, true
}
