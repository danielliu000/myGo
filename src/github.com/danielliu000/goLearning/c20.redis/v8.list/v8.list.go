package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/my/repo/src/github.com/danielliu000/goLearning/c20.redis/config"
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

func lPush(key string, values ...interface{}) {
	rs, err := rdb.LPush(ctx, key, values).Result()
	if err != nil {
		panic(err)
	} else if rs != 1 {
		fmt.Println(rs, " => 添加失败：")
	} else {
		fmt.Println(rs, " => 添加成功")
	}

}
func lRem(key string, count int64, value interface{}) {
	rs, err := rdb.LRem(ctx, key, count, value).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	}
	fmt.Printf("移除了%v个元素\n", rs)
}
func main() {
	lPush("user", "Daniel")
	lRem("user",2, "Leo")

}
