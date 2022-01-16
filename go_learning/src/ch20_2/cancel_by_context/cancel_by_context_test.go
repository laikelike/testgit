package cancelbycontext_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}

/*
Context与任务取消
当相关联的多个子任务存在，需要取消某一有子任务的任务时，需要连带者
取消其子任务，这时就可以用到context
根Context：通过context.Background()创建
子Context：context.WithCancel(parentContext)创建
	ctx,cancel := context.WithCancel(context.Background())
	把根节点传入之后，返回两个值前一个是context后面是cancel方法。
	cancel方法是用来取消当前调用的context值的，前一个context可
	以传到子任务里，使得子任务依次被取消
当前Context被取消时，基于他的子context都会被取消
接收取消通知 <-ctx.Done()

*/
