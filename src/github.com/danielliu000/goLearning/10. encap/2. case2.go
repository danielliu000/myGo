package main

import (
	"fmt"
	"model"
)

func main() {
	a := model.NewAccount()
	a.SetAcctNo("daniel000")
	a.Deposit(200)
	a.SetPassword("123123")
	fmt.Println(*a)
}
