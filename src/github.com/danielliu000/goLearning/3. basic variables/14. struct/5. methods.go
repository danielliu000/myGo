package main

import "fmt"

type Person1 struct {
	Name string  `json:"name"`
	Age byte			`json:"age"`
	Gender string `json:"gender"`
}
func (p Person1) test () {  // 只能被Person类型调用
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
//如果一个类型实现了String()方法，则fmt.Println会默认调用String()方法来输出
func (p *Person1) String() string {
	str := fmt.Sprintf("p.Name=[%v], p.Age=[%v]", p.Name, p.Age)
	return str
}
func main() {

	p := Person1{
		Name:   "Daniel",
		Age:    23,
		Gender: "male",
	}
	//p.test()
	fmt.Println(&p)
}
