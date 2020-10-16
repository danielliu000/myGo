package main

import (
	"fmt"
	"reflect"
)

//给一个变量 var v float64 = 1.2 请用反射得到它的Value,然后获取相应的Type, Kind值。
//并将reflect.Value转化为interface{},再将interface{}转化成 float64


func reflectFloat64(i interface{}) {
	rVal := reflect.ValueOf(i)
	rType := reflect.TypeOf(i)
	rKind := rVal.Kind()

	iv := rVal.Interface()
	num := iv.(float64)

	fmt.Printf("rVal=%v, rTpye=%v, rKind=%v, num=%v", rVal,rType,rKind,num)


}


func main() {
	var v float64 = 1.2
	reflectFloat64(v)
}
