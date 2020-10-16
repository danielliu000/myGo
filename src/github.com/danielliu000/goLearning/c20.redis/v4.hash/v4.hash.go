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

//向key中添加field元素的值
func hashSet(key, field string, data interface{}) int64 {
	val, err := rdb.HSet(ctx, key, field, data).Result()
	if err != nil {
		panic(err)
	}
	return val //设置成功返回1  否则返回0
}

//根据key 获取field的值
func hashGet(key, field string) string {
	val, err := rdb.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		fmt.Printf("key %v does not exist.\n", key)
	} else if err != nil {
		panic(err)
	}
	return val
}

//批量设置多个field的值
func batchHashSet(key string, fields map[string]interface{}) bool {
	val, err := rdb.HMSet(ctx, key, fields).Result()
	if err != nil {
		panic(err)
	}
	return val
}

//批量获取key中fields的值
func batchHashGet(key string, fields ...string) map[string]interface{} {
	resMap := make(map[string]interface{})
	val, err := rdb.HMGet(ctx, key, fields...).Result()
	if err != nil {
		panic(err)
	} else if err == redis.Nil {
		fmt.Printf("kesy does not exist\n")
	} else {
		for i, field := range fields {
			resMap[field] = val[i]
		}
	}
	return resMap

}

//
func main() {
	defer rdb.Close()
	//设置Hash的field值, 如果key没有该field则设置返回1, 否则返回0
	fmt.Println()
	fmt.Println("-----调用hashSet------")
	fmt.Println()
	rs1 := hashSet("user", "name", "Daniel")
	if rs1 == 1 {
		fmt.Println("设置成功")
	} else if rs1 == 0 {
		fmt.Println("该field已经存在, 设置失败")
	}

	//也可以不用返回值,无提示
	hashSet("user", "gender", "male")

	fmt.Println()
	fmt.Println("------调用hashGet-------")
	fmt.Println()
	rs2 := hashGet("user", "gender")
	fmt.Println(rs2)

	fmt.Println()
	fmt.Println("------调用batchHashSet-------")
	fmt.Println()
	userMap := make(map[string]interface{})
	userMap["name"] = "Daniel"
	userMap["age"] = 12
	userMap["state"] = "NSW"
	flag := batchHashSet("user", userMap)
	if flag {
		fmt.Println("批量设置成功")
	} else {
		fmt.Println("批量设置失败")
	}

	fmt.Println()
	fmt.Println("------调用batchHashGet-------")
	fmt.Println()
	for k, v := range batchHashGet("user", "name", "age", "state") {
		fmt.Printf("%v:%v\n", k, v)
	}

}
