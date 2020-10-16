package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.128.200:6379",
		Password: "",
		DB:       0, //default db 0 (0 - 15)
	})
	var ctx = context.Background()

	err := rdb.Set(ctx, "name", "Daniel", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

}