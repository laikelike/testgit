package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() { //类似于Java中finally中释放锁
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	//加等待时间是因为外面协程运行的速度超过了所有里面协程执行完的速度
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

//上一个方法强制等了一会拖延了时间，wait方法不用等
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() { //类似于Java中finally中释放锁
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done() //外面wait()等待完成的是这个
		}()
	}
	//等待所有协程结束才执行后面代码
	wg.Wait()
	t.Logf("counter = %d", counter)
}
