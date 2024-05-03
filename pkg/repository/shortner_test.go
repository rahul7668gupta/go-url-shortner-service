package repository_test

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
	"github.com/rahul7668gupta/go-url-shortner/pkg/logger"
	"github.com/rahul7668gupta/go-url-shortner/pkg/repository"
	"github.com/rahul7668gupta/go-url-shortner/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestShortnerRepo_CreateShortCodeRecord(t *testing.T) {
	shortCode := "abc"
	testUrl := "https://www.infracloud.io"
	ctx := context.Background()
	loggerObj := logger.InitLogger()
	redisClient, clientMock := redismock.NewClientMock()

	// set redis client mock expectations
	clientMock.ExpectSet(shortCode, testUrl, 0).SetVal(testUrl)

	// Initialise Repo
	shortnerRepo := repository.NewShortnerRepository(redisClient, loggerObj)
	// Call test func
	err := shortnerRepo.CreateShortCodeRecord(ctx, shortCode, testUrl)
	// Assert Errors
	assert.Nil(t, err)
}

func TestShortnerRepo_IncrementDomainCounter(t *testing.T) {
	testUrl := "https://www.infracloud.io"
	urlDomain, _ := utils.GetUrlDomain(testUrl)
	ctx := context.Background()
	loggerObj := logger.InitLogger()
	redisClient, clientMock := redismock.NewClientMock()

	// set redis client mock expectations
	clientMock.ExpectZIncrBy(constants.DOMAIN_COUNTER, 1, urlDomain).SetVal(1)

	// Initialise Repo
	shortnerRepo := repository.NewShortnerRepository(redisClient, loggerObj)
	// Call test func
	err := shortnerRepo.IncrementDomainCounter(ctx, urlDomain)
	// Assert Errors
	assert.Nil(t, err)
}

func TestShortnerRepo_CreateIndexOnOriginalUrl(t *testing.T) {
	shortCode := "abc"
	testUrl := "https://www.infracloud.io"
	ctx := context.Background()
	loggerObj := logger.InitLogger()
	redisClient, clientMock := redismock.NewClientMock()

	// set redis client mock expectations
	clientMock.ExpectSet("index:url:"+testUrl, shortCode, 0).SetVal(shortCode)

	// Initialise Repo
	shortnerRepo := repository.NewShortnerRepository(redisClient, loggerObj)
	// Call test func
	err := shortnerRepo.CreateIndexOnOriginalUrl(ctx, testUrl, shortCode)
	// Assert Errors
	assert.Nil(t, err)
}

func TestShortnerRepo_LookupURL(t *testing.T) {
	shortCode := "abc"
	testUrl := "https://www.infracloud.io"
	ctx := context.Background()
	loggerObj := logger.InitLogger()
	redisClient, clientMock := redismock.NewClientMock()

	// set redis client mock expectations
	clientMock.ExpectGet("index:url:" + testUrl).SetVal(shortCode)

	// Initialise Repo
	shortnerRepo := repository.NewShortnerRepository(redisClient, loggerObj)
	// Call test func
	actualShortCode, foundUrl := shortnerRepo.LookupURL(ctx, testUrl)
	// Assert Errors
	assert.Equal(t, true, foundUrl)
	assert.Equal(t, shortCode, actualShortCode)
}
