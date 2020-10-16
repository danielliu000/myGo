package main

import "fmt"

func main() {
	// 使用切片的三种方式

	//testing. 定义一个切片，然后去引用一个已经创建好的数组
	var arr = [...]int {1,2,3,4,5}
	var slice1 = arr[1:3]
	fmt.Println(slice1)
	//2. 通过make创建切片 此时切片必须make后才可使用
	var slice2 []int = make([]int, 4, 10) //长度为4 容量为10 类型为int的切片
	slice2[1] = 10
	slice2[3] = 20
	fmt.Println(slice2)
	//3. 定义时直接指定一个数组
	var slice3 = []string {"Tom", "Jerry" ,"Daniel" ,"Winnie"}
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))


 //方式1和方式2的区别
 //方式1的数组可见，可以直接操作数组 slice = mySort[x:y]指向 数组下标为x的地址
 //方式2slicemake时，会指向一个数组地址为下标0。而该数组是不可见的，即没有一个数组变量指向它，
 // 仅可以被slice操作

}
