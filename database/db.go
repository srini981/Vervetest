package database

import "github.com/go-redis/redis"

var RedisClient *redis.Client

func Initialize() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}
