package main

import "fmt"

func main() {
	var arr [5]float64
	fmt.Printf("mySort's addr: %p\n", &arr)
	for i := 0; i < 5; i++ {
		fmt.Println(arr[i]) //默认为0
		arr[i] = float64(i)
		fmt.Printf("mySort[%v]'s addr: %p\n", i, &arr[i])
		// 地址连续，每个地址占用类型字节数，这里float64占8字节
	}
	fmt.Println(arr)


}
