package main

import (
	"fmt"
	"reflect"
)

//结构体获取tag标签， 遍历字段值、修改字段值
//调用机构体方法（通过传地址方式完成）修改bp1代码


type Person2 struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score float32 `json:"score"`
	Gender string `json:"gender"`
}

func ( m *Person2) Print2() {
	fmt.Println("----Start-----")
	fmt.Println(*m)
	fmt.Println("------End------")
}

func (m *Person2) GetSum2(n1, n2 int) int {
	return n1 + n2
}

func(m *Person2) Set2(name string, age int, score float32, gender string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Gender = gender
}
func TestStruct2(i interface{}) {
	rVal := reflect.ValueOf(i)
	rType := reflect.TypeOf(i)
	rv := rVal.Elem()  //Person2 结构体
	fmt.Println(rv.Kind())  //struct
	fmt.Println(rType)  // *main.Person2  ptr
	numFields := rv.NumField()
	fmt.Println("字段个数:", numFields)

	//修改字段值
	rv.Field(0).SetString("Leo")
	rv.Field(1).SetInt(43)
	rv.Field(2).SetFloat(78.4)
	rv.Field(3).SetString("male")

	//打印标签、字段值
	for i := 0; i < numFields; i++ {
		tagVal := rType.Elem().Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("字段tag标签：%v, 字段值： %v\n",tagVal,rv.Field(i) )
		}
	}
	//调用Print方法
	rVal.Method(1).Call(nil)

	//调用GetSum方法
	var args []reflect.Value
	args = append(args, reflect.ValueOf(15))
	args = append(args, reflect.ValueOf(30))
	res1 := rVal.Method(0).Call(args)
	fmt.Println("res = ", res1[0].Int())

	//调用Set方法
	var fieldArgs []reflect.Value
	fieldArgs = append(fieldArgs, reflect.ValueOf("Winnie"))
	fieldArgs = append(fieldArgs, reflect.ValueOf(32))
	fieldArgs = append(fieldArgs, reflect.ValueOf(float32(98.2)))
	fieldArgs = append(fieldArgs, reflect.ValueOf("female"))
	rVal.Method(2).Call(fieldArgs)
	rVal.Method(1).Call(nil)
}
func main() {
	person2 := Person2{
		Name:   "Daniel",
		Age:    33,
		Score:  84,
		Gender: "male",
	}
	TestStruct2(&person2)
}
