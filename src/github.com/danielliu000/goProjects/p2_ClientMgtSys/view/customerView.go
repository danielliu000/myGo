package view

import (
	"fmt"
	"github.com/danielliu000/goProjects/p2_ClientMgtSys/model"
	"github.com/danielliu000/goProjects/p2_ClientMgtSys/service"
	"strings"
)

type CustomerView struct {
	key             int
	customerService service.CustomerService
}

func (c *CustomerView) MainMenu() {

Loop:
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println(" testing 添加客户")
		fmt.Println(" 2 修改客户")
		fmt.Println(" 3 删除客户")
		fmt.Println(" 4 客户列表")
		fmt.Println(" 5 退出")
		fmt.Print("请选择(testing-5)：")
		_, _ = fmt.Scanln(&c.key)
		switch c.key {
		case 1:
			c.Add()
		case 2:
			c.Edit()
		case 3:
			c.Delete()
		case 4:
			c.Show()
		case 5:
			if c.Exit() {
				break Loop
			}
		default:
			fmt.Println("你的输入有误，请重新输入...")

		}
	}
	fmt.Println("你退出了客户关系管理系统...")
}
func (c *CustomerView) Show() {
	fmt.Println("编号\t 姓名\t\t性别\t\t年龄\t\t    电话\t\t      邮箱")
	customers := c.customerService.GetCustomersList()
	for _, v := range customers {
		fmt.Println(v.GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

func (c *CustomerView) Add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Print("姓名:")
	name := ""
	_, _ = fmt.Scanln(&name)
	fmt.Print("性别:")
	gender := ""
	_, _ = fmt.Scanln(&gender)
	fmt.Print("年龄:")
	age := 0
	_, _ = fmt.Scanln(&age)
	fmt.Print("电话:")
	phone := ""
	_, _ = fmt.Scanln(&phone)
	fmt.Print("电邮:")
	email := ""
	_, _ = fmt.Scanln(&email)
	customer := model.NewCustomer(name, gender,age, phone, email)
	if c.customerService.AddCustomer(*customer) {
		fmt.Println("---------------------添加完成---------------------")
	}else {
		fmt.Println("---------------------添加失败---------------------")
	}
}
func (c *CustomerView) Delete() {
	id := -1
	fmt.Println("---------------------删除客户---------------------")
	fmt.Printf("请选择待删除客户的编号(-1退出): ")
	_, _ = fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Printf("确认是否删除(y/n): ")
	keyPress := ""
	_, _= fmt.Scanln(&keyPress)
	if strings.ToLower(keyPress) == "y" {
		if c.customerService.DeleteCustomer(id) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("--------------删除失败, 输入的ID不存在-------------")
		}
	}else if strings.ToLower(keyPress) == "n" {
		return
	}

}
func (c *CustomerView) Edit() {
	id := -1
	customers := c.customerService.Customers

	fmt.Println("---------------------修改客户---------------------")
	fmt.Printf("请选择待修改客户编号（-1退出): ")
	_, _ = fmt.Scanln(&id)
	index := c.customerService.FindById(id)
	if id == -1 {
		return
	} else if index >= 0 && index < len(customers) {
		fmt.Printf("姓名（%v）: ", customers[index].Name)
		name := customers[index].Name
		_, _ = fmt.Scanln(&name)
		fmt.Printf("性别(%v): ", customers[index].Gender)
		gender := customers[index].Gender
		_, _ = fmt.Scanln(&gender)
		fmt.Printf("年龄(%v): ", customers[index].Age)
		age := customers[index].Age
		_, _ = fmt.Scanln(&age)
		fmt.Printf("电话(%v): ", customers[index].Phone)
		phone := customers[index].Phone
		_, _ = fmt.Scanln(&phone)
		fmt.Printf("电邮(%v):", customers[index].Email)
		email := customers[index].Email
		_, _ = fmt.Scanln(&email)

		c.customerService.EditCustomer(id, name, gender, age,phone,email)

		fmt.Println("---------------------修改完成---------------------")
	} else {
		fmt.Println("输入的编号不正确")
	}

}
func (c *CustomerView) Exit() bool {
	keyPress := ""
	for {
		fmt.Printf("你确定要退出吗? y/n: ")
		_, _ = fmt.Scanln(&keyPress)
		if strings.ToLower(keyPress) == "y" {
			fmt.Println()
			fmt.Println("程序退出, 再见")
			return true
		} else if strings.ToLower(keyPress) == "n" {
			return false
		} else {
			fmt.Println("请输入正确的指令: y or n")
		}
	}

}