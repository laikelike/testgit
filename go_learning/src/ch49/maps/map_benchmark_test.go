package maps

import (
	"strconv"
	"sync"
	"testing"
)

const (
	NumOfReader = 90
	NumOfWriter = 10
)

type Map interface {
	Set(key interface{}, val interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}

func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < NumOfWriter; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Set(strconv.Itoa(i), i*i)
					hm.Set(strconv.Itoa(i), i*i)
					hm.Del(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		for i := 0; i < NumOfReader; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Get(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncmap(b *testing.B) {
	b.Run("map with RWLock", func(b *testing.B) {
		hm := CreateRWLockMap()
		benchmarkMap(b, hm)
	})

	b.Run("sync.map", func(b *testing.B) {
		hm := CreateSyncMapBenchmarkAdapter()
		benchmarkMap(b, hm)
	})

	b.Run("concurrent map", func(b *testing.B) {
		superman := CreateConcurrentMapBenchmarkAdapter(199)
		benchmarkMap(b, superman)
	})
}

/*
性能比较：
读写次数相同：
enchmarkSyncmap/map_with_RWLock-8   193           5845559 ns/op
BenchmarkSyncmap/sync.map-8         100          10285535 ns/op
BenchmarkSyncmap/concurrent_map-8   546           2275829 ns/op

写是读的二倍：
enchmarkSyncmap/map_with_RWLock-8   102          11367627 ns/op
BenchmarkSyncmap/sync.map-8         62           19559811 ns/op
BenchmarkSyncmap/concurrent_map-8   300           3857359 ns/op

读是写的二倍：
enchmarkSyncmap/map_with_RWLock-8   166           7138456 ns/op
BenchmarkSyncmap/sync.map-8         100          11278795 ns/op
BenchmarkSyncmap/concurrent_map-8   416           3129653 ns/op

90%的写10%的读：
enchmarkSyncmap/map_with_RWLock-8   775           1364955 ns/op
BenchmarkSyncmap/sync.map-8         913           1315436 ns/op
BenchmarkSyncmap/concurrent_map-8   1927            616665 ns/op

10%的写90%的读
enchmarkSyncmap/map_with_RWLock-8                   753           1395961 ns/op
BenchmarkSyncmap/sync.map-8                          904           1337740 ns/op
BenchmarkSyncmap/concurrent_map-8                   1944            616631 ns/op
*/

/*
别让性能被“锁”住
1.减少锁的影响范围
2.减少发生锁冲突的概率，如sync.map或ConcurrentMap
3.避免使用锁：
高性能数据交换队列：LAMX Disruptor https://martinfowler.com/articles/lmax.html

*/
