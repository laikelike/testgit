package lock_test

import (
	"fmt"
	"sync"
	"testing"
)

var cache map[string]string

const NUM_OF_READER int = 40
const READ_TIMES = 100000

func init() {
	cache = make(map[string]string)

	cache["a"] = "aa"
	cache["b"] = "bb"
}

func lockFreeAccess() {
	var wg sync.WaitGroup
	wg.Add(NUM_OF_READER)
	for i := 0; i < NUM_OF_READER; i++ {
		go func() {
			for j := 0; j < READ_TIMES; j++ {
				_, err := cache["a"]
				if !err {
					fmt.Println("Nothing!")
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func lockAccess() {
	var wg sync.WaitGroup
	wg.Add(NUM_OF_READER)
	m := new(sync.RWMutex)
	for i := 0; i < NUM_OF_READER; i++ {
		go func() {
			for j := 0; j < READ_TIMES; j++ {
				m.RLock()
				_, err := cache["a"]
				if !err {
					fmt.Println("Nothing")
				}
				m.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkLockFree(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lockFreeAccess()
	}
}

func BenchmarkLock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lockAccess()
	}
}

/*
go test -bench=.

enchmarkLockFree-8          260           4536542 ns/op
BenchmarkLock-8                7         145937114 ns/op

go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof
top -cum

flat  flat%   sum%        cum   cum%
0.12s  0.65%  0.65%      9.62s 52.45%  WorkSpace/go_learning/src/ch49/lock_test.lockAccess.func1
1.46s  7.96%  8.62%      7.52s 41.00%  WorkSpace/go_learning/src/ch49/lock_test.lockFreeAccess.func1
5.22s 28.46% 37.08%      6.30s 34.35%  runtime.mapaccess2_faststr
4.69s 25.57% 62.65%      4.70s 25.63%  sync.(*RWMutex).RLock (inline)
4.55s 24.81% 87.46%      4.56s 24.86%  sync.(*RWMutex).RUnlock (inline)

list lockAccess

30ms       30ms     44:                   for j := 0; j < READ_TIMES; j++ {
.         4.70s     45:                           m.RLock()
50ms      290ms     46:                           _, err := cache["a"]
40ms       40ms     47:                           if !err {


结论：读写锁也会影响读的性能，不要乱加锁


*/

/*
sync.map
适合读多写少，且Key相对稳定的环境
采用了空间换时间的方案，并且采用指针的方式间接实现值的映射，所以存储空间会较built-in map大

分了两块区域：只读区和读写区(脏区),两个区域存储的指针指向同一块value区
只读区可以在无锁的情况下访问，新增和修改的数据在脏区
当在只读区找不到数据时，就会miss掉然后再脏区找(多读了一次只读区，比RW锁更糟了)
当多次miss后数据就会从脏区加入到只读区

*/

/*
Concurrent Map(蔡超实现的)
原理：缩小锁的范围，降低锁冲突的概率，提高读写速度
适用于读写都很频繁的操作
*/
