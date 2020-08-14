package main

import "fmt"

func main() {
	//const

	const (
		a = 4
		b = 2
		str = `hello, world`
	)
	c := a * b
	fmt.Println(c)
	fmt.Printf("%T:  %[1]v\n", str)
}
