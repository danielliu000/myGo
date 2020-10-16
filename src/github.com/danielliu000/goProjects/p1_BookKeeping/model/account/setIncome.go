package account

import "fmt"

func (a *account) setIncome() bool {
	moneyInput := 0.0
	fmt.Printf("本次收入金额：")
	_, _ = fmt.Scanln(&moneyInput)
	if moneyInput >= 0 {
		a.Amount = moneyInput
		a.Balance += moneyInput
		return true
	} else {
		fmt.Println("输入金额有误")
		return false
	}
}