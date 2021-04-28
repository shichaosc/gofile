package cache

import "github.com/go-redis/redis"

var Rd *redis.Client

func InitRedis() {
	Rd = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func init() {
	InitRedis()
}
