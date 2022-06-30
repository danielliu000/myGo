package b

import "fmt"

// 必须大写, 否则跨包无法访问, init 除外

type User struct {
	Name string
	Age  int
	Addr string
}

func init() {
	fmt.Println("This is an init func from package b")
}

func (u *User) getName() string {
	return u.Name
}
