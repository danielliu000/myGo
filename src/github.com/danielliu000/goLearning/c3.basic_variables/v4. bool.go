package main

import "fmt"

func main() {
	//bool 只有 true and false 占用1个字节
	a := bToi(true)
	b := iTob(3)
	fmt.Println(a,b)
}

//bool to int

func bToi (bool b) int {
	if b {
		return 1
	}
	return 0
}
//int to bool
func iTob (int i) bool { return i != 0 }

