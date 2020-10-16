package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

//find redis.conf:
//注释 bind 127.0.0.1 bind
//protect-mode no
//配置firewall-cmd --permanent --add-port=6379/tcp
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.128.200:6379",
	Password:"",
		DB:0,
})
	var ctx = context.Background()
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)

}
