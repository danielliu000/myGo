package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCpu := runtime.NumCPU()
	fmt.Println(numCpu)
	runtime.GOMAXPROCS(numCpu - 1) //Golangv1.8后 默认多核运行，不需要再手动设置

}
