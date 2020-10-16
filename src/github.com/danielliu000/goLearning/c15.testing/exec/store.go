package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

const FILEPATH = "E:/Daniel Files/IT_Learning/mygo/" +
	"src/github.com/danielliu000/goLearning/c15.testing/exec/person.json"

func (p *Person) Store() error {
	jsonData, err := json.Marshal(&p)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ioutil.WriteFile(FILEPATH, jsonData, 0666)
	if err != nil {
		fmt.Println("write file error: ", err)
		return err
	}
	fmt.Println("序列化保存成功")
	return nil

}
func (p *Person) ReStore() error {
	data, err := ioutil.ReadFile(FILEPATH)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(data, p)
	if err != nil {
		fmt.Println("unmarshal file error: ", err)
		return err
	}
	fmt.Println("反序列化读取成功")
	return nil
}

//func main() {
//	p := Person{
//		Name:   "Daniel",
//		Age:    33,
//		Gender: "male",
//	}
//	p.Store()
//	p.ReStore()
//}
