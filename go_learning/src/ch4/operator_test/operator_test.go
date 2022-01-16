package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4} //数组有三个点，切片没有三个点
	b := [...]int{1, 2, 2, 4}
	d := [...]int{1, 2, 3, 4}
	//与java不同，go可进行数组比较(相同维数，相同长度，且内容、顺序都相同)
	t.Log(a == b)
	t.Log(a == d)
}

const (
	Readable  = 1 << iota //0001
	Writeable             //0010
	Excutable             //0100
)

//&^按位置零
// a &^ b
//若b为1，a为任何值，结果都为0
//若b为0，a为何值，结果就为何值
func TestConstantTry1(t *testing.T) {
	a := 7            //0111
	a = a &^ Readable //0111 &^ 0001 = 0110
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Excutable == Excutable)
}
