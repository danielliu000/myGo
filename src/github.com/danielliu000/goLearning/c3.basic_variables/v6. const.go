package main

import "fmt"

func main() {
	//const

	const (
		a = iota   // 0
		//b, c, d 依次 +1
		b = iota
		c   //iota可以省略
		d
		e, f = iota, iota  // 同一行值不会递增  e=4, f=4
		str = `hello, world`
	)
	num := d * c
	fmt.Printf("a=%v, b=%v, c=%v, d=%v\n", a, b ,c, d)
	fmt.Printf("e=%v, f=%v\n",e, f)
	fmt.Println(num)
	fmt.Printf("str's type=%T, str's value=%v\n", str, str)
}
