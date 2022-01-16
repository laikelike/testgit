package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

//快速定义常量方式
const (
	Readable  = 1 << iota //0001
	Writeable             //0010
	Excutable             //0100
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 //0111
	t.Log(a&Readable == Readable, a&Writeable, a&Excutable)
}
