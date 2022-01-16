package autogrowing

import "testing"

const NumOfElems = 100000
const times = 1000

func TestAutoGrow(t *testing.T) {
	for i := 0; i < times; i++ {
		s := []int{}
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestProperInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, 100000)
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestOverSizeInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, 800000)
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAutoGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{}
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkProperInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, NumOfElems)
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkOverSizeInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, NumOfElems*8)
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

/*
避免内存分配和复制注意：
1.初始化至合适的大小，因为自动扩容是有代价的
2.复用内存

测试数据：
BenchmarkAutoGrow-8        1617    811581 ns/op
BenchmarkProperInit-8      6600    161025 ns/op
BenchmarkOverSizeInit-8    1338    970501 ns/op


*/
