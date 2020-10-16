package main

import "fmt"

func main() {
	//create
	studentsMap := make(map[string]map[string]string, 2) //长度2 可以动态扩容

	studentsMap["stu01"] = map[string]string{
		"name":   "Daniel",
		"gender": "male",
	}
	studentsMap["stu02"] = map[string]string{
		"name":   "Winnie",
		"gender": "female",
	}
	//可以动态增加了
	studentsMap["stu03"] = map[string]string{
		"name":   "Mia",
		"gender": "female",
	}
	fmt.Println(studentsMap)
	//update
	studentsMap["stu02"]["name"] = "Lucy"
	fmt.Println(studentsMap)

	//remove--- delete
	//delete(studentsMap,"stu01")
	//fmt.Println(studentsMap)
	//一次性删除所有key
	//testing. 遍历keys, 依次删除
	//2. 重新make 使原来的map没有被引用，被gc回收
	//studentsMap = map[string]map[string]string{}
	//fmt.Println(studentsMap)
	//
	//find
	val, isThere := studentsMap["stu03"]
	if isThere {
		fmt.Println("yes", val["name"])
	} else {
		fmt.Println("no" )
	}
}

