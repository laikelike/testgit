package sync_pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}
	v := pool.Get().(int) //第一次取出一定会创建新的对象
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
	//创建的对象并不会自动的放在对象池里面，所以get的又是新对象
	v2, _ := pool.Get().(int)
	fmt.Println(v2)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
sync.Pool总结

1.适用于通过复用，降低复杂对象的创建和GC代价
2.是协程安全的，所以会有锁的开销
3.生命周期受GC影响，不适合用于做连接池等，需自己管理生命周期的资源的池化
4.更像是对象的缓存。

对于sync.pool每个Processor分为两部分：私有对象和共享池
私有对象区只能放一个对象，是协程安全的；共享池是协程不安全的。

sync.Pool对象的获取
每次从对象池获取对象时，首先从私有对象中找，如若没有会从当前协程的
Processer中的共享池中获取，有返回，没有就从其他的Processor的共享池
中获取，如果都没有，就用sync.Pool指定的new函数产生一个新的对象返回。
sync.Pool对象的放回
首先考虑放回到私有对象，若私有对象无则保存为私有对象
若私有对象存在，放入当前Processor子池的共享池中
sync.Pool对象的生命周期
GC会清除sync.pool缓存的对象
对象的缓存有效期为下一次GC之前，我们不能人为控制GC时间，所以无法把控
sync.pool对象的存活时间，也就不能把它当对象池用

Processer:不是真正的处理器，是go语言实现的协程处理器每个processor上都
挂有准备运行的协程队列，processor依次运行这些协程
*/
