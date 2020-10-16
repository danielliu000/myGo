package main

import "fmt"

func main() {
	type Cat struct {
		Name string
		Age int
		Color string
	}

	var cat1 Cat
	cat1.Name = "xiaohua"
	cat1.Age = 2
	cat1.Color = "black"
	fmt.Println(cat1)

	var cat2 Cat
	cat2.Name = "xiaoye"
	cat2.Age = 3
	cat2.Color = "white"
	fmt.Println(cat2)
}
