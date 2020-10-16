package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

//结构体序列化
func testStruct() {
	person := Person{
		Id:       1,
		Name:     "Daniel",
		Age:      33,
		Gender:   "male",
		Birthday: "1000-03-23",
	}
	jsonData, err := json.Marshal(&person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}
//Map序列化
func testMap() {
	var myMap map[string]interface{}
	myMap = make(map[string]interface{})
	myMap["name"] = "Daniel"
	myMap["age"] = 33
	myMap["gender"] = "male"
	myMap["birthday"] = "2121-12-12"

	jsonData, err := json.Marshal(&myMap)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}
//切片序列化
func testSlice() {
	var mySlice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "Tom"
	m1["age"] = 15
	m1["gender"] = "female"


	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "jack"
	m2["age"] = 5
	m2["gender"] = "male"
	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	jsonData, err := json.Marshal(mySlice)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

}
//普通数据类型序列化
func testFloat64 () {
	num1 := 33.3
	jsonData, err := json.Marshal(num1)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

}



func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
