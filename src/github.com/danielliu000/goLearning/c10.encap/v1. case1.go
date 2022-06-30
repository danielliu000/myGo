package main

import (
	"fmt"
	"model/person" // GoPath 下 找到Model/person
)

func main() {
	p := person.NewPerson("Daniel")
	p.SetAge(30)
	p.SetSalary(30000)
	fmt.Println(*p)
	fmt.Println(p.GetAge())
	fmt.Println(p.GetSalary())
}
