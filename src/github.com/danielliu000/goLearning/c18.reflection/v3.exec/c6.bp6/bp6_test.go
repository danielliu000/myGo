package test

import (
	"reflect"
	"testing"
)

//使用反射创建并操作结构体

type user struct {
	UserId string
	Name string
}

func TestReflectStructPtr(t *testing.T) {
	var (
		model *user
		st reflect.Type
		elem reflect.Value
	)

	st = reflect.TypeOf(model) //获取类型 *user
	t.Log("reflect.TypeOf", st.Kind().String()) //ptr
	st = st.Elem()   //st指向的类型
	t.Log("reflect.TypeOf.Elem", st.Kind().String()) //struct

	//New 返回 一个Value类型值，该值持有一个指定类型type的新申请的零值指针
	elem = reflect.New(st)
	t.Log("reflect.New", elem.Kind().String())  //ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String()) //struct

	//model是创建的user结构体变量
	model = elem.Interface().(*user) //类型断言 model是*user，它指向和elem是一样的
	elem = elem.Elem() //取得指向的值
	elem.FieldByName("UserId").SetString("123456")
	elem.FieldByName("Name").SetString("nickName")
	t.Log("model model.name", model, model.Name)  //通过elem改变后 model也被改变 都指向*user，

}