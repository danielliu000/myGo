package main

import "fmt"

func main() {

	//for...range
	var myArr = [...]string {0:"daniel",1:"winnie",2:"Leo",3:"Mia"}
	for idx, val := range myArr {
		fmt.Printf("%v --- %v\n", idx, val)
	}
}
