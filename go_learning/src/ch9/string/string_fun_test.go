package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",") //用","分割字符串
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-")) //用"-"连接字符串
}

//字符串与其他类型的转换
func TestConv(t *testing.T) {
	s := strconv.Itoa(10) //int类型转字符串 10-->"10"
	t.Log("str " + s)
	//字符型转int型会返回两个返回值，需要if判断，i接受返回结果，err接受错误
	if i, err := strconv.Atoi("20"); err == nil {
		t.Log(10 + i)
	} else {
		t.Log("convert is failed!") //若为汉字则转换失败
	}
	//t.Log(10 + strconv.Atoi("10"))
}
