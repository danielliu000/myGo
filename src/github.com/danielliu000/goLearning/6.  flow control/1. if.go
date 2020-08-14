package main

import "fmt"

func main() {
	var ageInput byte
	fmt.Println("input your age: ")
	_, err := fmt.Scanln(&ageInput)

	if (err != nil) { //可以加括号,不报错 但强烈不推荐
		fmt.Println("Error: ", err)
	}

	if age := ageInput; age > 18 { //支持if作用域内声明变量,此时不能加小括号
		fmt.Println("ok, enter") //即时只有一条语句， 也必须有大括号
	} else { // else 不能换行
		fmt.Println("refused")
	}
}

