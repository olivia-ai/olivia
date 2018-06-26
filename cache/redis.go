package cache

import (
	"github.com/go-redis/redis"
	"os"
)

var (
	RedisAddress  = os.Getenv("REDIS_ADRESS")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
)

// Returns a new cache client with default database
func CreateClient() redis.Client {
	address := "localhost:6379"
	if RedisAddress != "" {
		address = RedisAddress
	}

	password := ""
	if RedisPassword != "" {
		password = RedisPassword
	}

	return *redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
}
