package select_test

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	time.Sleep(time.Millisecond * 500)
	return "Service is Done"
}
func AsyncService() chan string {

	//retCh := make(chan string) //用make声明一个channel
	retCh := make(chan string, 1)
	go func() {
		ret := Service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("Service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {

	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}

/*
select多路选择机制：
select {
case ret := <-retCh1:
	t.Logf("result %s",ret)
case ret := <- retCh2 :
	t.Logf("result %s",ret)
default:
	t.Error("No one returned")
}
只要有任何一个协程结束就会执行此case所定义的部分，
如果都没准备好就走default
*/
