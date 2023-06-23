package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default db
	})

	// Context
	ctx := context.Background()

	// Ping the Redis server to check the connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis, err: ", err)
		return
	}
	fmt.Println("Successfully connected to Redis")

	// Store data
	err = client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		log.Fatal("Failed to insert data in redis, err: ", err)
	}
	fmt.Println("Successfully stored data in Redis")

	// Get data
	value, err := client.Get(ctx, "foo").Result()
	if err != nil {
		log.Fatal("Failed to get data from redis, err: ", err)
	}
	fmt.Println("Value: ", value)
}

// TODO env vars ADDR
