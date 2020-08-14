package main

//声明自定义全局类型
type myIntType int32
type sumFuncType func(myIntType, myIntType) int


func getSum(a, b myIntType) int {
	return int(a + b)
}
//func applySum(sumFunc func(int, int) int, x int, y int ) int { //声明过长
//	return sumFunc(x, y)
//}
func applySum(sumFunc sumFuncType, x, y myIntType ) int { //应用自定义函数类型
	return sumFunc(x, y)
}

func main() {
	//自定义数据类型
	type myInt int
	// go认为自定义类型是独立类型
	// 即：myInt 和 int 不是一个类型
	var a myInt = 30
	var b int
	a = a + myInt(b) // 直接 a + b 会报错
	println("a=", a)

	applySum(getSum, 30, 40)
}

