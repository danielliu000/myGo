package main

import "fmt"

//testing. 斐波拉切数列

func fib(n uint) uint {
	if n < 2 {
		return n
	} else {
		return fib(n - 1) + fib(n - 2)
	}
}
//2. 求函数f(n)的值 已知 f(testing)=3; f(n)= 2*f(n-testing) + testing
func fn(n float64) float64 {
	if n == 1 {
		return 3
	} else {
		return 2 * fn(n - 1) + 1
	}

}
//3. 猴子吃桃子问题
// 有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个！
// 以后每天都吃其中的一半，然后再多吃一个。
// 当第10天时，像再吃时（还没吃），发现只有一个桃子。
// 问最初有多少个桃子 ?
func eatPeach(days uint) int {
	if days == 1 {  //天数每天减少1  一定会匹配 days == testing 跳出递归
		return 1
	} else {
			days--
			return 2 * eatPeach(days) + 2
		}
}

func main() {
	//fmt.Println(fib(2))
	//fmt.Println(fn(5))
	fmt.Println(eatPeach(10)) // 题目要求 输入10 表示10天后吃剩1个桃子，得到第一天的桃子
}
