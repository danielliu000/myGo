package account

import "fmt"

func (a *account) getDetail() {
	if a.Flag {
		fmt.Println("-----------------当前收支明细记录-----------------")
		fmt.Println(a.Title)
		a.Detail += fmt.Sprintf("%v\t\t%v\t\t%v\t\t%v\n", a.IncomeExpense, a.Balance, a.Amount, a.Desc)
		fmt.Printf(a.Detail)
	} else {
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
}