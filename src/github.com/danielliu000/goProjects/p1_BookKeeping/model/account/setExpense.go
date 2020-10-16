package account

import "fmt"

func (a *account) setExpense() bool {
	moneyInput := 0.0
	fmt.Printf("本次支出金额：")
	_, _ = fmt.Scanln(&moneyInput)
	if moneyInput >= 0 && moneyInput <= a.Balance {
		a.Amount = moneyInput
		a.Balance -= moneyInput
		return true
	} else {
		fmt.Println("输入金额有误")
		return false
	}
}