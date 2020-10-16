package service

//Customers CRUD operation
import (
	"fmt"
	"github.com/danielliu000/goProjects/p2_ClientMgtSys/model"
)

type CustomerService struct {
	Customers   []model.Customer
	CustomerNum int
}

func NewCustomerService() *CustomerService {
	customerService := CustomerService{}
	customerService.CustomerNum = 1
	customer := model.NewCustomer("",
		"", 0, "", "")
	customerService.Customers = append(customerService.Customers, *customer)
	return &customerService
}

func (c *CustomerService) GetCustomersList() []model.Customer {
	return c.Customers
}

func (c *CustomerService) AddCustomer(customer model.Customer) bool {
	c.CustomerNum++
	customer.Id = c.CustomerNum
	c.Customers = append(c.Customers, customer)
	return true
}

func (c *CustomerService) FindById(id int) int {
	index := -1
	for i, v := range c.Customers {
		if v.Id == id {
			fmt.Println("V.ID: ", v.Id)
			index = i
		}
	}
	return index
}

func (c *CustomerService) EditCustomer(id int, name string,
	gender string, age int, phone string, email string) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	} else {
		c.Customers[index] = model.Customer{
			Id:     id,
			Name:   name,
			Gender: gender,
			Age:    age,
			Phone:  phone,
			Email:  email,
		}
		//c.Customers[index].Name = name
		//c.Customers[index].Gender = gender
		//c.Customers[index].Age = age
		//c.Customers[index].Phone = phone
		//c.Customers[index].Email = email
		return true
	}
}

func (c *CustomerService) DeleteCustomer(id int) bool {
	index := c.FindById(id)
	fmt.Println("index value: ", index)
	if index == -1 {
		return false
	} else {

		c.Customers = append(c.Customers[:index], c.Customers[index+1:]...)
		//for i := index; i < len(c.Customers); i++ {
		//	c.CustomerNum = c.Customers[i].Id - testing
		//	c.Customers[i].Id = c.CustomerNum
		//}
		return true
	}
}
