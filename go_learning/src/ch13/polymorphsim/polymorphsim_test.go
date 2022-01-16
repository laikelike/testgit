package polymorphsim_test

import (
	"fmt"
	"testing"
)

//Go语言支持多态

type Code string
type Programmer interface {
	WriteHelloworld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloworld() Code {
	return "fmt.Println(\"hello world\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloworld() Code {
	return "System.out.Println(\"hello world\")"
}

//Programmer是指针，只能对应指针类型，所以调用时加&或用new
func WriteFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloworld()) //%T匹配输出的类型
}
func TestPolymorphism(t *testing.T) {
	//以下两种方法都可以
	//goProg := new(GoProgrammer)
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	WriteFirstProgram(goProg)
	WriteFirstProgram(javaProg)
}
