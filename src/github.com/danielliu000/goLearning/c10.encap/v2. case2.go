package main

//需要项目打开go.mod所在目录, 包才可见

import (
	"encap/model/a"   //引入 gomod中的名字(encap)/包路径(model/a)
	_ "encap/model/b" //不引用包中变量, 但init 在引用加载时 会执行

	"fmt"
)

func main() {

	sum := a.Add(3, 5) //调用 包名.成员变量
	fmt.Println(sum)

}

func init() {
	fmt.Println("This is an init func from main") //最后执行
}
