package database

import (
	"context"
	"fmt"
	"go-api/pkg"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() {
	// fmt.Println("Connecting to Redis...")
	host := pkg.GetEnv("REDIS_HOST", "localhost")
	port := pkg.GetEnv("REDIS_PORT", "6379")

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		DB:       0,
		Password: pkg.GetEnv("REDIS_PASSWORD", ""),
	})

	redisCtx := context.Background()
	if _, err := client.Ping(redisCtx).Result(); err != nil {
		panic("Failed to connect to Redis")
	}
	fmt.Println("Connected to Redis")
}
func GetRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", pkg.GetEnv("REDIS_HOST", "localhost"), pkg.GetEnv("REDIS_PORT", "6379")),
	})
}
func CloseRedisClient(client *redis.Client) {
	err := client.Close()
	if err != nil {
		fmt.Println("Error closing Redis client:", err)
	} else {
		fmt.Println("Redis client closed")
	}
}
