package account

import "fmt"

func (a *account) setDesc(label string) {
	desc := ""
	fmt.Printf(label)
	_, _ = fmt.Scanln(&desc)
	a.Desc = desc
}