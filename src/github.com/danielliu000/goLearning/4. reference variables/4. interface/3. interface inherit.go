package main

type CInterface interface {
	testC()
}

type DInterface interface {
	testD()
}


//如果要实现E接口，需要把B,C 和自身E 接口的方法都实现
type EInterface interface {
	CInterface
	DInterface
	testE()
}


type Person struct {

}


func (p *Person) testC() {

}
func (p *Person) testD() {

}
func (p *Person) testE() {

}




func main() {
	p := Person{}
	var a EInterface = &p

	a.testC()
	a.testD()
	a.testE()
}
