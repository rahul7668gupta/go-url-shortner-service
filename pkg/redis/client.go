package redis

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
	"github.com/redis/go-redis/v9"
)

func InitRedisClient() (*redis.Client, error) {
	host := os.Getenv(constants.REDIS_ENDPOINT)
	if len(host) == 0 {
		panic("unable to read REDIS_ENDPOINT from env")
	}
	password := os.Getenv(constants.REDIS_PASSWORD)
	db := os.Getenv(constants.REDIS_DB)
	dbInt, err := strconv.Atoi(db)
	if err != nil {
		dbInt = 0
	}

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       dbInt,
	})

	_, pingErr := client.Ping(context.TODO()).Result()
	if pingErr != nil {
		log.Panic("redis server not reachable,error=", pingErr)
	}
	return client, nil
}
