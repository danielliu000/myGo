package main

import "testing"

func TestPerson_Store(t *testing.T) {
	p := Person{
		Name:   "Daniel",
		Age:    33,
		Gender: "male",
	}
	res := p.Store()
	if res != nil {
		t.Fatalf("执行错误, 希望值=%v, 实际发生错误=%v", nil, res)
	}
	t.Logf("执行成功")

}

func TestPerson_ReStore(t *testing.T) {
	p := Person{}
	res := p.ReStore()
	if res != nil {
		t.Fatalf("执行错误, 希望值=%v, 实际发生错误=%v", nil, res)
	}
	if p.Name != "Daniel" {
		t.Fatalf("执行错误, 值不对 期望=%v 实际=%v", "Daniel", p.Name)
	}

	t.Logf("执行成功")

}
