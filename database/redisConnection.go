package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func RedisInstance() *redis.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	RedisURL := os.Getenv("REDIS_URL")

	client := redis.NewClient(&redis.Options{
		Addr:     RedisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Redis!")

	return client
}

var RedisClient *redis.Client = RedisInstance()
