package person

import "fmt"

type person struct {
	Name string
	age int
	salary float64
}

//工厂函数
func NewPerson(name string) *person  {
	return &person{
		Name: name,
	}
}
//为了访问age， salary 需要一对get和set函数

func (p *person) GetAge() int {
	return p.age
}
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	}else {
		fmt.Println("输入年龄不正确")
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}
func (p *person) SetSalary(salary float64) {
	if salary > 3000 && salary < 40000 {
		p.salary = salary
	}else {
		fmt.Println("输入薪水不正确")
	}
}

