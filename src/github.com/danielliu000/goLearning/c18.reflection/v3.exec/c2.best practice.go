package main

import (
	"fmt"
	"reflect"
)

//1. 使用反射：
//遍历结构体字段，
//调用结构体方法，
//并获取结构体标签的值
type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score float32 `json:"score"`
	Gender string `json:"gender"`
}

func ( m Person) Print() {
	fmt.Println("----Start-----")
	fmt.Println(m)
	fmt.Println("------End------")
}

func (m Person) GetSum(n1, n2 int) int {
	return n1 + n2
}

func(m Person) Set(name string, age int, score float32, gender string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Gender = gender
}
func TestStruct(i interface{}) {
	rType := reflect.TypeOf(i)
	rVal := reflect.ValueOf(i)
	rKind := rVal.Kind()

	num := rVal.NumField() //结构体字段数量
	fmt.Printf("strunct has %v fileds\n", num)

	if rKind != reflect.Struct {
		fmt.Println("expect a struct")
		return
	}

	for i := 0; i < num; i++ {
		fieldName := rType.Field(i).Name    //type.Filed() 返回的是一个 StructField结构体
		fieldValue := rVal.Field(i)         //value.Field() 返回的是一个 Value结构体
		fmt.Printf("Person结构体 %v 字段的值： %v\n", fieldName, fieldValue)
		fieldTag :=rType.Field(i).Tag.Get("json")
		if fieldTag != "" {
			fmt.Printf("Person结构体 %v Tag的值: %v\n",fieldName, fieldTag)
		}

	}
	numMethods := rVal.NumMethod()
	fmt.Println("结构体的方法个数：", numMethods)
	//调用Print()方法
	rVal.Method(1).Call(nil)  //方法排序是按函数名的ASCII码来排序 GetSum(0), Print(1), Set(2)

	//调用GetSum
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(30))
	res := rVal.Method(0).Call(params)
	fmt.Printf("res = %v", res[0].Int())
}
func main() {
	person := Person{
		Name:   "Daniel",
		Age:    33,
		Score:  84,
		Gender: "male",
	}
	TestStruct(person)
}
