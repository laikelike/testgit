package type_test

import (
	"math"
	"testing"
)

type Myint int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	var c Myint
	b = int64(a) //必须要显示类型转换
	c = Myint(b)
	//b = a 不支持隐式类型转换
	t.Log(a, b, c)
	d := math.MaxInt32
	t.Log(d)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a //取地址a
	//指针类型不可以进行运算，所以也就不可以进行地址自加1,连续访问数组
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr) //%T 获得变量类型
}

func TestString(t *testing.T) {
	var s string //默认初始化字符串为空值
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" { //若字符串s未被赋值
		s = "laikekkkk"
		t.Log("*" + s + "*")
	}
}
