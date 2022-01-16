package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

/*
go func(){}()代表一个协程
协程是更轻量级的线程，只有2k大小(线程1M)
协程和系统线程(kernal Space Entity)对应关系是多对多
Java是一对一
*/
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ { //这里是值传递
		go func(i int) {
			fmt.Println(i)
		}(i) //把i传入这个用协程跑的函数中，复制了一份i传入
		//每个协程的i相互不影响

		//错误示范,i是被共享的，需要使用锁机制
		// go func() {
		// 	fmt.Println(i)
		// }()
	}
	time.Sleep(time.Microsecond * 50)
}

/*
Processer:不是真正的处理器，是go语言实现的协程处理器每个processor上都
挂有准备运行的协程队列，processor依次运行这些协程

提高协程并发的两种机制：
1.当某一协程长时间占用Processor时，Go机制中会有一个守护线程记录每个
processor完成的协程数量，若某个processor长时间没动时，会往协程的
任务栈中插入一个标记，当此协程运行时遇到非内联函数时，会读到这个标记
并将自己中断，插到等候协程队列队尾，切换别的协程继续运行。
2.当某一协程被系统中断了(如等待IO)，processor会把自己移到另一个系统
线程中，继续执行其他协程，当被中断的协程被唤醒，会把自己加到某个Processor
的协程队列中，或全局等待队列当中。
*/
