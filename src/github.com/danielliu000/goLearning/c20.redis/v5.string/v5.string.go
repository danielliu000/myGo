package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/my/repo/src/github.com/danielliu000/goLearning/c20.redis/config"
	"time"
)

var rdb *redis.Client
var ctx = context.Background()

//初始化 连接 redis DB
func init() {
	rdb = redis.NewClient(&redis.Options{
		Password:   config.Password,
		DB:         config.DB,
		Addr:       config.RedisAddr,
		MaxRetries: 3,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		//fmt.Println(pong)
		fmt.Println("数据库连接成功")
	}
}

//SetNX 已有key不会覆盖
func stringSetNX(key string, value interface{}, exp time.Duration) bool {
	err := rdb.SetNX(ctx, key, value, exp).Err()
	if err != nil {
		panic(err)
	}
	return true
}

func stringGet(key string){
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	}
	fmt.Printf("%v: %v\n",key, val)
}

func main() {
	stringSetNX("name", "Leo", time.Second * 60)
	stringGet("name")
}
