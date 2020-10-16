package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//练习：
//1. 创建一个Person结构体【name, age, address】
//2. 使用rand方法配合随机创建10个Person实例，并放入到channel中
//3. 遍历channel， 将各个Person实例信息显示在终端

type Person struct {
	Name string
	Age int
	Address string
}
func main() {

	seed := time.Now().UnixNano()
	rand.Seed(seed)
	personChan := make(chan Person, 10)
	for i := 1; i <= 10; i++ {
		personChan <- Person {
			Name: "Daniel_" + strconv.Itoa(rand.Intn(100)),
			Age: rand.Intn(100),
			Address: "Sydney_" + strconv.Itoa(rand.Intn(100)),
		}
	}
	for j := 1; j <= 10; j++ {
		person := <- personChan
		fmt.Printf("person'Name: %v, person's Age: %v, person's Address: %v\n",
			person.Name, person.Age, person.Address)
	}


}
