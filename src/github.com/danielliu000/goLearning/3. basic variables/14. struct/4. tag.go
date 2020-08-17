package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name string  `json:"name"`
		Age byte			`json:"age"`
		Gender string `json:"gender"`
	}

	p1 := Person{
		Name:   "Daniel",
		Age:    30,
		Gender: "male",
	}
	//序列化json字符串
	data, _ := json.Marshal(p1)
	fmt.Println(string(data)) //没有tag时， Jason keys 都是大写，不符合规范
}
