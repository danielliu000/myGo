package main

import "fmt"

func main() {
	type Person struct {
		Age int
		Name string
	}

	var p1 Person
	p1.Age = 10
	p1.Name = "Daniel"

	var p2 *Person = &p1
	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)
	p2.Name = "tom~"
	fmt.Printf("p2.Name= %v, p1.Nmae=%v\n",p2.Name, p1.Name)
	fmt.Printf("p2.Name= %v, p1.Nmae=%v\n",(*p2).Name, p1.Name)

}
