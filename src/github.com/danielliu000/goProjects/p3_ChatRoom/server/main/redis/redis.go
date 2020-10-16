package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	Rdb *redis.Client
)

func InitRedisClient() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:"192.168.128.200:6379",
		DB: 0,
		Password:"",
		MaxRetries: 3,
	})
	ctx := context.Background()
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		//fmt.Println(pong)
		fmt.Println("数据库连接成功")
	}

	
}