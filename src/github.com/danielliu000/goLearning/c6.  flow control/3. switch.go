package main

import "fmt"

func main() {
	//golang switch
	//testing. 没有break,
	//2. case 可以跟多个表达式(变量，常量，有返回值的函数..) 以逗号分隔
	//3. default 不是必须的
	//4. switch 后 的条件类型要与case后的类型匹配
	//5. case 后表达式是常量的话，不能相同
	//6. switch后可以不带表达式 case当成 if-else分支使用
	//7. switch后可以直接声明变量 以分号结尾
	//8. 关键字fallthrough 类似continue
	n := 10
	m := 20
	switch { //switch后不跟表达式

	case n >= 10:
		fmt.Println(n)
		fallthrough //继续switch 不会终止
	case m >= 20:
		fmt.Println(20)
	default:
		fmt.Println("not equal to 10")
	}

	switch x := 20; { //直接声明变量
	case x >= 20:
		fmt.Println(x)
	default:
		fmt.Println("not equal to 10")
	}

}