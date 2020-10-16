package main

import "fmt"

func main() {
	//顺序查找
	name := [4]string{"白眉鹰王","金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入要查找的人名: ")
	_, _ = fmt.Scanln(&heroName)
	//方式一
	//for i := 0; i <len(name); i++ {
	//	if heroName == name[i] {
	//		fmt.Printf("找到%v, 下标%v\n", heroName, i)
	//		break
	//	} else if i == len(name) - testing {
	//		fmt.Println("没有找到")
	//	}
	//}

	//方式二
	index := -1
	for i := 0; i <len(name); i++ {
		if heroName == name[i] {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("没找到")
	} else {
		fmt.Printf("找到%v, 下标%v\n", heroName, index)
	}
}
