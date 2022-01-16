package panicrecover_test

import (
	"errors"
	"fmt"
	"testing"
)

//panic用于不可恢复的错误
//panic退出前会执行defer指定的内容
//os.Exit()退出时不会调用defer指定的函数
//os.Exit()退出时不输出当前调用栈信息
func TestPanicVxExit(t *testing.T) {
	defer func() { //用os.Exit(-1)退出不会执行此函数
		fmt.Println("Finally!")
	}()
	fmt.Println("start")
	//os.Exit(-1)
	panic(errors.New("Something wrong!"))
}
