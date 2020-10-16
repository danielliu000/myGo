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

func zSet(key string, members *redis.Z) {
	rs, err := rdb.ZAdd(ctx, key, members).Result()
	if err != nil {
		panic(err)
	} else if rs != 1 {
		fmt.Println(rs, " => 设置失败：")
	} else {
		fmt.Println(rs, " => 设置成功")
	}

}
func zGet(key string, start, stop int64) {
	val, err := rdb.ZRange(ctx, key, start, stop).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
func zRem(key string, members ...interface{}) {
	rs, err := rdb.ZRem(ctx, key, members).Result()
	if err != nil {
		panic(err)
	} else if rs != 1 {
		fmt.Println(rs, " => 移除失败：")
	} else {
		fmt.Println(rs, " => 移除成功")
	}

}
func main() {
	var scores redis.Z
	scores.Score = 80
	scores.Member = "Daniel"
	zSet("scores", &scores)
	scores.Score = 90
	scores.Member = "Winnie"
	zSet("scores", &scores)

	zGet("scores", 0 , -1)
	zRem("scores", "Daniel")
}
