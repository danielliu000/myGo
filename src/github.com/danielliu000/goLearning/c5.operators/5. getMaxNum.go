package main

import "fmt"

//get larger number from 2 numbers
func getLarge(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//get larger number from 3 numbers
func getMax(a int, b int, c int) int {
	if c >= getLarge(a, b) {
		return c
	} else {
		return getLarge(a, b)
	}
}
func main() {

	fmt.Println(getLarge(15, 25))
	fmt.Println(getMax(3, 6, 3))
}