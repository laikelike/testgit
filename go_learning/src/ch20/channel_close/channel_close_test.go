package channelclose_test

import (
	"fmt"
	"sync"
	"testing"
)

//生产者消费者模型

func dataProducer(ch chan int, wg *sync.WaitGroup) { //生产者
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) //关闭channel,不可以再往  通道送数据
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) { //消费者
	go func() {
		//for i := 0; i < 10; i++ {
		for { //有了通道关闭机制，就不需要知道应该接受多少数据了
			//ok为bool值，true表示正常接收，false表示通道关闭
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg) //若有多个生产者，也不影响
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}

/*
当消费者并不知道有多少生产者有多少数据可放时，需要一种机制告诉消费
者，数据放完了。
思路一：可以用约定好的标记来表明数据放完了(如往通道中放-1)，但此时
需要知道有多少个消费者，才能知道放多少个-1，这里消费者和生产者数量
的耦合度就很高了
思路二：channel有关闭机制，当生产者生产结束，没有数据可放时，关闭
channel，关闭后再往通道放数据会返回两个值：
数据值：返回通道类型的0值.布尔值：true表示正常接收，false表示通道关闭
此时再往通道传值，bool值返回false,会异常退出，若不判断bool值，会返回
通道类型0值。


向关闭的channel发送数据，会导致panic
v, ok <- ch ; ok为bool值，true表示正常接受，false表示通道关闭
所有channel接受者都会再channel关闭时，立刻从阻塞等待中返回且上述
ok值为false。这个广播机制常被利用，向多个订阅者同时发送信号，
如：退出信号
*/
