package main

import "fmt"

func main() {
	studentsMap := make(map[string]map[string]string)

	studentsMap["stu01"] = map[string]string{
		"name":   "Daniel",
		"gender": "male",
	}
	studentsMap["stu02"] = map[string]string{
		"name":   "Winnie",
		"gender": "female",
	}

	for k1, v1 := range studentsMap{
		fmt.Printf("studnet %v: ", k1)
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}

}
