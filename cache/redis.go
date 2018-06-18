package cache

import "github.com/go-redis/redis"

// Returns a new cache client with default database
func CreateClient(address string, password string) redis.Client {
	return *redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
}
