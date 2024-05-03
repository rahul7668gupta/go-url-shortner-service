package repository_test

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
	"github.com/rahul7668gupta/go-url-shortner/pkg/dto"
	"github.com/rahul7668gupta/go-url-shortner/pkg/logger"
	"github.com/rahul7668gupta/go-url-shortner/pkg/repository"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestShortnerRepo_GetMetrics(t *testing.T) {
	ctx := context.Background()
	loggerObj := logger.InitLogger()
	redisClient, clientMock := redismock.NewClientMock()
	redisMetricsData := []redis.Z{
		{
			Member: "https://www.infracloud.io",
			Score:  5,
		},
		{
			Member: "https://www.udemy.com",
			Score:  3,
		},
		{
			Member: "https://www.amazon.com",
			Score:  1,
		},
	}
	expectedMetrics := []dto.Metrics{
		{
			DomainName: "https://www.infracloud.io",
			Count:      5,
		},
		{
			DomainName: "https://www.udemy.com",
			Count:      3,
		},
		{
			DomainName: "https://www.amazon.com",
			Count:      1,
		},
	}

	// set redis client mock expectations
	clientMock.ExpectZRevRangeWithScores(constants.DOMAIN_COUNTER, 0, 2).SetVal(redisMetricsData)

	// Initialise Repo
	shortnerRepo := repository.NewShortnerRepository(redisClient, loggerObj)
	// Call test func
	metrics, err := shortnerRepo.GetMetrics(ctx)
	// Assert Errors
	assert.Nil(t, err)
	assert.Len(t, metrics, 3)
	assert.Equal(t, expectedMetrics, metrics)
}
