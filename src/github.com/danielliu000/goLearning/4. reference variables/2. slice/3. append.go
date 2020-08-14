package main

import "fmt"

func main() {
	arr := []string {"Tom", "Jerry" ,"Daniel" ,"Winnie"}
	slice1 := arr[:]

	// append 增加元素
	slice2 := append(slice1, "Leo","Mia")
	fmt.Println(slice1) // slice3本身不会发生变化
	fmt.Println(slice2)

	//还可以追加一个切片 以...方式
	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)


	//注意: append方法后指向了单独的数组,是值拷贝,不是引用,因此不会改变任何原始切片或数组
	fmt.Println("------append赋值, 值拷贝-------")
	slice1[0] = "Ryan"
	// 此时 slice1 指向 mySort, 而slice2, slice3使用过了append 分别指向的是新的不可见数组
	fmt.Println("slice1=", slice1)  //通过赋值改变slice1指向的arr[0]的元素
	fmt.Println("slice2=", slice2)   //不变
	fmt.Println("slice3=", slice3) // 不变
	fmt.Println("mySort=", arr) // 不变

	fmt.Println("-------slice赋值, 引用--------")
	slice1[0] = "Michael"
	fmt.Println("slice1=", slice1)
	fmt.Println("slice2=", slice2) //不变
	fmt.Println("slice3=", slice3) //不变
	fmt.Println("mySort=", arr) // 变了

}
