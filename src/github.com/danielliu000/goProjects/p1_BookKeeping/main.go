package main

import (
	"github.com/danielliu000/goProjects/p1_BookKeeping/model/account"
)

func main() {
	acct := *account.NewAccount()
	account.Menu(&acct)
}
