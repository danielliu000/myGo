package main

import (
	"fmt"
	"reflect"
)

//作业
//编写一个Cal结构体， 又两个字段Num1, Num2
//方法 GetSub(name, string)
//使用返回遍历Cal结构体所有的字段信息
//使用反射机制完成对GetSub的调用，输出形式为
//“tom 完成了减法运行， 8 - 3 = 5”

type Cal struct {
	Num1 int
	Num2 int
}

func (c *Cal) GetSub(name string) string {
	return fmt.Sprintf("%v 完成了减法运行", name)
}

func reflectCal (i interface{}) {
	rv := reflect.ValueOf(i)
	elem := rv.Elem()
	rk := elem.Kind()
	num := elem.NumField()
	if rk != reflect.Struct {
		fmt.Println("Expect Struct")
		return
	}
	nums := make([]int, num)
	for i := 0; i < num; i++ {
		nums[i] = elem.Field(i).Interface().(int)
	}

	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf("Tom")
	res := rv.Method(0).Call(args)

	sub := nums[0] - nums[1]
	fmt.Printf("%v, %v - %v = %v", res[0].String(),nums[0], nums[1], sub )
}

func main() {
	cal := Cal{
		Num1: 8,
		Num2: 3,
	}
	reflectCal(&cal)
}
