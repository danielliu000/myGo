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

func setsSet(key string, members ...interface{}) {
	rs, err := rdb.SAdd(ctx, key, members).Result()
	if err != nil {
		panic(err)
	} else if rs != 1 {
		fmt.Println(rs, " => 设置失败：")
	} else {
		fmt.Println(rs, " => 设置成功")
	}
}

func setsGet(key string) {
	rs, err := rdb.SMembers(ctx, key).Result()  //Redis `SMEMBERS key` command output as a slice.
	//rs, err := rdb.SMembersMap(ctx, key).Result()  //Redis `SMEMBERS key` command output as a map.
	if err != nil {
		panic(err)
	}
	fmt.Println(rs)
}

func setsRem(key string, members ...interface{}) {
	rs, err := rdb.SRem(ctx, key, members).Result()
	if err != nil {
		panic(err)
	} else if rs != 1 {
		fmt.Println(rs, " => 移除失败：")
	} else {
		fmt.Println(rs, " => 移除成功")
	}

}

func main() {
	setsSet("mySet", "Sydney")
	setsGet("myset")
	setsRem("mySet","XIAN")
}
