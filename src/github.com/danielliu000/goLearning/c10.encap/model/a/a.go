package a

import "fmt"

func Add(a, b int) int {
	return a + b
}

func init() {
	fmt.Println("This is an init func from package a")
}
