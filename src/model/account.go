package model

import "fmt"

type account struct {
	accountNo string  // len: 6-10
	balance float64    // > 20
	password string    //len : 6
}

func NewAccount () *account {
	return &account{
		accountNo: "",
		balance:    0,
		password:   "",
	}
}

func (a *account) GetAcctNo() string {
	return a.accountNo
}
func (a *account) GetBalance() float64 {
	return a.balance
}
func (a *account) GetPassword() string {
	return a.password
}

func (a *account) SetAcctNo(acctNo string) {
	if len(acctNo) >= 6 && len(acctNo) <= 10 {
		a.accountNo = acctNo
	} else {
		fmt.Println("账号设置不正确")
	}
}
func (a *account) Deposit(money float64)  {
	if money > 20 {
		a.balance += money
	} else {
		fmt.Println("存款金额不正确")
	}
}
func (a *account) SetPassword(password string) {
	if len(password) == 6 {
		a.password = password
	} else {
		fmt.Println("密码设置不正确")
	}
}


