package panicrecover_test

import (
	"errors"
	"fmt"
	"testing"
)

//当心recover后，形成僵尸服务进程，导致health check失效
//有时候让进程重启释放资源反而有利于程序执行
//例如如下recover并没有任何有效操作，不过是记录了错误而已
func TestPanic(t *testing.T) {
	defer func() {
		//recover接收err并处理
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("Somthing wrong!"))
	//os.Exit(-1)
	//fmt.Println("End")

}
