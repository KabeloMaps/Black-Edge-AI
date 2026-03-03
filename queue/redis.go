package queue

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var RDB *redis.Client

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("❌ Redis connection failed:", err)
	}
	log.Println("✅ Redis connected")
}