package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.128.200:6379",
		Password:"",
		DB: 0,
	})
	ctx := context.Background()
	get := rdb.Do(ctx, "get", "name" )
	if err := get.Err(); err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}

	val, _ := get.Text()  //shortcut
	fmt.Println(val)  //Daniel

	fmt.Println(get.String()) //get name: Daniel

}
