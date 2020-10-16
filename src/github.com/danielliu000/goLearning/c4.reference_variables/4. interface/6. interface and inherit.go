package main

import "fmt"

//利用接口扩充结构体
type Monkey struct {
	Name string
}

type BirdAbility interface {
	flying()
}

type FishAbility interface {
	swimming()
}
type LittleMonkeyAbility interface {
	BirdAbility
	FishAbility
}
func (m Monkey) climbing () {
	fmt.Printf("%v can always climb\n", m.Name)
}

type LittleMonkey struct {
	Monkey
}
//LittleMonkey实现fish 和 bird的接口 达到扩展的目的 同时不破坏继承
func (m LittleMonkey) flying() {
	fmt.Printf("%v can fly now\n",m.Name)
}
func (m LittleMonkey) swimming() {
	fmt.Printf("%v can swim now\n", m.Name)
}
func main() {
	lm := LittleMonkey{Monkey{Name:"LittleMonkey"}}
	lm.climbing()
	var la LittleMonkeyAbility
	la = lm
	la.swimming()
	la.flying()


}
