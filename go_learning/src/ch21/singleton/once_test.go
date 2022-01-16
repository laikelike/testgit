package singleton_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj)) //%x：16进制
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
线程安全的单例模式
sync.Once里有个DO()方法给DO()方法里传入一个函数，确保在多线程
的环境下只调用一次
*/
