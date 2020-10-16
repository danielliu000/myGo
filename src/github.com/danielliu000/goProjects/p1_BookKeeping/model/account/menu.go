package account
import "fmt"

func Menu(account *account) {
Loop:
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println(" testing 收支明细")
		fmt.Println(" 2 登记收入")
		fmt.Println(" 3 登记支出")
		fmt.Println(" 4 退出软件")
		fmt.Print("请选择(testing-4)：")
		_, _ = fmt.Scanln(&account.Key)

		switch account.Key {
		case 1:
			account.getDetail()
		case 2:
			fmt.Println("-----------------登记收入-----------------")
			account.IsDone = account.setIncome()

			if account.IsDone {
				account.setDesc("本次收入说明: ")
				account.setIncomeExpense("收入")
				account.Flag = true
				fmt.Println("录入成功")

			} else {
				fmt.Println("录入失败")
			}

		case 3:
			fmt.Println("-----------------登记支出-----------------")
			account.IsDone = account.setExpense()
			if account.IsDone {
				account.setDesc("本次支出说明: ")
				account.setIncomeExpense("支出")
				fmt.Println("录入成功")
				account.Flag = true
			} else {
				fmt.Println("录入失败")
			}

		case 4:
			if account.Exit() {
				break Loop
			}

		default:
			fmt.Println("您输入的选项不正确")
		}
	}

}