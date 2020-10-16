package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//实现对Hero结构体切片的排序sort.Sort(data Interface)

// testing. 声明Hero 结构体

type Hero struct {
	Name string
	Age int
}

//2. 声明Hero结构体切片
type HeroSlice []Hero

//3. 实现interface 接口 Len, Less, Swap
func (hs HeroSlice) Len() int {
	return len(hs)
}

//Less决定使用什么标准排序 升序/降序
//按Hero的Age从小到大排序
func (hs HeroSlice) Less(i, j int) bool {

	//按Hero的Age从小到大排序
	//return hs[i].Age < hs[j].Age
	//按Hero的Name从大到小排序
	return hs[i].Name > hs[j].Name
}


func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
func main() {
	var heroes  HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	//排序前
	for _, v := range heroes{
		fmt.Println(v)
	}

	sort.Sort(heroes)
	fmt.Println("-------排序后----------")
	for _, v := range heroes{
		fmt.Println(v)
	}
}
