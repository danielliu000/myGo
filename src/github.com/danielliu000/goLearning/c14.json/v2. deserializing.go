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

const (
	jsonData1 = "{\"id\":testing,\"name\":\"Daniel\"," +
		"\"age\":33,\"gender\":\"male\",\"birthday\":\"1000-03-23\"}"

	jsonData2 = "[{\"age\":15,\"gender\":\"female\",\"name\":\"Tom\"}," +
		"{\"age\":5,\"gender\":\"male\",\"name\":\"jack\"}]"
)
func unMarshalStruct () {
	person := Person{}
	err := json.Unmarshal([]byte(jsonData1), &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person)
}
func unMarshalMap () {

	//反序列化时不需要再make了
	var myMap map[string]interface{}

	err := json.Unmarshal([]byte(jsonData1), &myMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(myMap)
}
func unMarshalSlice () {
	var mySlice []map[string]interface{}

	err := json.Unmarshal([]byte(jsonData2), &mySlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mySlice)
}
func main() {
	unMarshalStruct()
	unMarshalMap()
	unMarshalSlice()
}
