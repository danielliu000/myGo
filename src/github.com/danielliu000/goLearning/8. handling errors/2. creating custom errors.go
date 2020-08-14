package main

import (
	"errors"
	"fmt"
)

func readConf (name string) (err error) { //返回error类型的错误
	if name == "config.conf" {
		return nil
	} else {
		return errors.New("读取失败") //自定义错误
	}
}
func checkRead(name string) {
	err := readConf(name)
	if err != nil {
		//发生错误 需要抛出错误 终止程序
		panic(err)
	} else {
		fmt.Println("读取成功 程序继续执行")
	}
}

func main() {
	checkRead("config.conf")
	checkRead("config.ini") //失败后,程序终止执行
	checkRead("config.conf")

}
