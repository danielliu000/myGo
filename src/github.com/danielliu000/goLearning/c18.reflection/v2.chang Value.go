package main

import (
	"fmt"
	"reflect"
)

func reflect01 (i interface{}) {
	rVal := reflect.ValueOf(i)
	rKind := rVal.Kind()   //ptr
	fmt.Println(rKind)
	rVal.Elem().SetInt(20)  //Elem()方法返回指针指向的值，SetInt() 改值
}
func main() {
	num := 10
	reflect01(&num)
	fmt.Println(num) //20
}
