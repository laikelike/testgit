package string_test

import (
	"testing"
	"unsafe"
)

//Unicode是一种字符编码方式，UTF-8是Unicode的存储实现
func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化默认零值
	s = "hello"
	t.Log(len(s))
	//s[1] = '3'  //string是不可变的byte slice
	s = "\xE4\xB8\xA5"     //可以存储任何二进制数据
	s = "\xE4\xBA\xB5\xFF" //任意给s赋一个二进制值

	s = "中"
	t.Log(s)
	t.Log(len(s)) //是byte的长度

	//rune是一种新的数据类型，可以取出字符串代表的Unicode值
	//rune等同于int32，常用来处理unicode或UTF-8字符
	c := []rune(s)
	t.Log(len(c))
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s) //"中"字在string/byte中存储:[0xE4,0xB8,0xAD]
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		//[1]的意思是%c和%d都对应着c这一个变量
		t.Logf("%[1]c %[1]d %[1]x", c) //%x输出的是rune而不是byte
	}
}
