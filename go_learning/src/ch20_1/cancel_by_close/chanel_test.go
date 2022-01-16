package cancelbyclose_test

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

//两种取消机制
//机制1：向通道传递几个值证明取消了几个协程
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
	//cancelChan <- struct{}{}
}

//机制2：关闭通道，松耦合，推荐做法
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}) //0：管道容量
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}
