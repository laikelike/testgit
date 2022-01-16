package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	//变量声明方法1,多用于全局或外部变量声明
	// var a int = 1
	// var b int = 1
	//变量声明方法2
	// var (
	// 	a int = 1
	// 	b  = 1
	// )
	//变量声明3，可进行自动类型推断
	a := 1
	b := 1
	//fmt.Print(a)
	t.Log(" ", a)
	for i := 0; i < 5; i++ {
		//fmt.Print(" ", b)
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	//fmt.Println()
}
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//交换方式1
	// tmp := a
	// a = b
	// b = tmp
	//交换方式2,可在一个赋值语句中对多个变量进行赋值
	a, b = b, a
	t.Log(a, b)
}
