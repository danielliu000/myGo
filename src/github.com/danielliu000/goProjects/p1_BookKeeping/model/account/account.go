package account

type account struct {
	IncomeExpense string  //收支  收入/支出
	Balance       float64 //账户金额
	Amount        float64 //收支金额
	Desc          string  //说明
	Key           int
	Title         string
	Detail        string
	Flag          bool
	IsDone        bool
}

func NewAccount() *account {
	return &account{
		IncomeExpense: "",
		Balance:       10000.0,
		Amount:        0.0,
		Desc:          "",
		Key : 0,
		Title:"收支\t 账户金额\t 收支金额\t说   明",
		Detail : "",
		Flag : false,
		IsDone : false,
	}
}






