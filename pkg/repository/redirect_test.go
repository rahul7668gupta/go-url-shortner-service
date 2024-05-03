package repository_test

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/rahul7668gupta/go-url-shortner/pkg/logger"
	"github.com/rahul7668gupta/go-url-shortner/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestShortnerRepo_GetOriginalUrlForShortCode(t *testing.T) {
	shortCode := "abc"
	testUrl := "https://www.infracloud.io"
	ctx := context.Background()
	loggerObj := logger.InitLogger()
	redisClient, clientMock := redismock.NewClientMock()

	// set redis client mock expectations
	clientMock.ExpectGet(shortCode).SetVal(testUrl)

	// Initialise Repo
	shortnerRepo := repository.NewShortnerRepository(redisClient, loggerObj)
	// Call test func
	originalUrl, err := shortnerRepo.GetOriginalUrlForShortCode(ctx, shortCode)
	// Assert Errors
	assert.Nil(t, err)
	assert.Equal(t, testUrl, originalUrl)
}
