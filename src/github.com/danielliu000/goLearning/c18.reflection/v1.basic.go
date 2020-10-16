package main

import (
	"fmt"
	"reflect"
)

//

type Student struct {
	Name string
	Age  int
}

func reflectTest01(b interface{}) {
	//通过反射获取传入变量的type kind 值
	//1. 先获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	//虽然Typeof, ValueOf 打印的是 int, 10, 但其实类型不是int
	//不能简单进行运算
	//n2 := rVal + 10  error
	fmt.Printf("rVal's Type=%T\n", rVal)

	n2 := rVal.Int() + 10 //需要调用方法
	fmt.Println(n2)

	//将rVal 转为interface{}
	iv := rVal.Interface()
	num := iv.(int) //用断言转回int类型

	fmt.Println(num)

}
func reflectTest02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)   // main.Student

	rVal := reflect.ValueOf(b)
	iv := rVal.Interface()
	fmt.Printf("%T\n", iv)

	stu, ok := iv.(Student)
	if ok {
		fmt.Println(stu.Name)
		fmt.Println(stu.Age)
	}
	//获取 Kind
	fmt.Println(rTyp.Kind())  // 方式一：struct
	fmt.Println(rVal.Kind())  // 方式二：struct

	//kind 范围比 type 大  struct  > main.Student
}
func main() {
	//演示对基本类型 interface{} reflect.Value进行反射的基本操作
	//num := 10
	//reflectTest01(num)

	stu := Student{
		Name: "Daniel",
		Age:  33,
	}
	reflectTest02(stu)
}
