package customertype

import (
	"fmt"
	"testing"
	"time"
)

//自定义类型
type IntConv func(op int) int

//专门用来计时的函数
func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n) //调用内部函数
		//输出内部函数运行的时间
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
