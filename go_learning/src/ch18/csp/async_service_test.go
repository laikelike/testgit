package csp

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	time.Sleep(time.Millisecond * 50)
	return "Service is Done"
}

func otherTask() {
	fmt.Println("Other task working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Other task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(Service())
	otherTask()
}

func AsyncService() chan string {

	//retCh := make(chan string) //用make声明一个channel
	//buffer channel，指定了buffer的容量大小
	retCh := make(chan string, 1)
	go func() { //用另外一个协程运行service，而不阻塞当前协程
		ret := Service()
		fmt.Println("returned result.")
		retCh <- ret //运行结束将结果放入channel
		fmt.Println("Service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {

	retCh := AsyncService() //另起炉灶，不影响otherTask协程的运行
	otherTask()
	fmt.Println(<-retCh) //从channel中取数据

}

/*
channel是有容量限制并独立处理Groutine的
channel有两种机制
1.通信的双方都必须在channel上才能完成交互，一方不在的时候另一方都
会被阻塞在那里.(未指定buffer大小，需要手把手送到)
2.为通道指定容量，只要通道未满都可以放，只要通道有数据都可以取。
机制一：
新建协程运行AsyncService()，主协程运行otherTask(),
0豪秒处两协程同时运行，输出Other task working on something else
50毫秒处service()运行结束，输出returned result.并将结果传入channel
此时Service()所在协程阻塞，等待另一协程取出
100毫秒处otherTask()执行结束，输出Other task is done.
主协程从channel中取数据并输出Service is Done，释放service()协程
最后输出Service exited.
机制二：
新建协程运行AsyncService()，主协程运行otherTask(),
0豪秒处两协程同时运行，输出Other task working on something else
50毫秒处service()运行结束，输出returned result.并将结果传入通道
传入结束，service()所在协程退出输出Service exited.
100毫秒处otherTask()执行结束，输出Other task is done.
最后需要时，输出channel中的数据Service is Done。
*/
