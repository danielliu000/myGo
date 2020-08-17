package person

type student struct {  //student是小写
	Name string
	Age int
	Gender string
}

//使用工长函数解决小写结构体无法导出问题
func NewStudent (name string, age int, gender string) *student {
	return &student{
		Name:   name,
		Age:    age,
		Gender: gender,
	}
}