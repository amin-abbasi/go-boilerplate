package services

import (
	"context"
	"log"
	"time"

	"github.com/amin-abbasi/go-boilerplate/configs"
	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)

// ConnectRedis initializes the Redis client.
func ConnectRedis() {
	host := configs.GetEnvVariable("REDIS_HOST")
	port := configs.GetEnvVariable("REDIS_PORT")
	addr := host + ":" + port
	log.Printf(">>> Redis Address: %v", addr)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("<<< Connected to Redis >>>")
}

// SetToken sets a token with expiration in Redis.
func SetToken(token string, expiration time.Duration) error {
	err := RedisClient.Set(context.Background(), token, "", expiration).Err()
	if err != nil {
		log.Printf("Failed to set token in Redis: %v", err)
		return err
	}
	return nil
}

// GetToken checks if a token exists in Redis.
func GetToken(token string) (bool, error) {
	result, err := RedisClient.Exists(context.Background(), token).Result()
	if err != nil {
		log.Printf("Failed to get token from Redis: %v", err)
		return false, err
	}
	return result == 1, nil
}

// DeleteToken deletes a token from Redis.
func DeleteToken(token string) error {
	err := RedisClient.Del(context.Background(), token).Err()
	if err != nil {
		log.Printf("Failed to delete token from Redis: %v", err)
		return err
	}
	return nil
}
