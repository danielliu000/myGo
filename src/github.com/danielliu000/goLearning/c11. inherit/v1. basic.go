package main

import "fmt"

type student struct {
	Name string
	Age int
	Score float64
}

func (stu *student) showInfo() {
	fmt.Printf("学生姓名：%v, 年龄： %v, 成绩： %v\n", stu.Name, stu.Age, stu.Score)
}
func (stu *student) setScore(score float64) {
	stu.Score = score
}


type Pupil struct {
	student //嵌入匿名结构体实现继承
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试...")
}

type Graduate struct {
	student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试...")
}


func main() {
	//继承使用方法
	p := Pupil{student{Score:0,Age:10,Name:"Daniel"}}
	g := Graduate{student{Score:79,Age:18,Name:"Winnie"}}

	p.testing()
	p.setScore(80)
	p.showInfo()

	g.testing()
	g.setScore(90)
	g.showInfo()
}
