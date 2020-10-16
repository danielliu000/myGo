//testing. 写一个程序，获取一个int变量num的地址，并显示
//2. 将num的地址赋值给指针 ptr，并通过ptr修改num的值

package main

import "fmt"

func main() {
	num := 10
	numAddr := &num
	fmt.Println("address of num : ", numAddr) //testing

	var ptr *int = numAddr
	*ptr = 20
	fmt.Println("value of num is : ", num)

}
