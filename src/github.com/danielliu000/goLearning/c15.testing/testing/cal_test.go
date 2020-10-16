package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	res := add(10)
	if res != 55 {
		//fmt.Println("")
		t.Fatalf("执行错误 期望值是%v, 实际值是%v", 55, res)
	}
	//如果正确，输出日志
	t.Logf("执行正确")
}

func TestGetSub(t *testing.T) {
	res := getSub(10,3)
	if res != 7 {
		//fmt.Println("")
		t.Fatalf("执行错误 期望值是%v, 实际值是%v", 7, res)
	}
	//如果正确，输出日志
	t.Logf("执行正确")
}