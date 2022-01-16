package passingref

import (
	"testing"
)

const NumOfElems = 1000

type Content struct {
	Detail [10000]int
}

// 传数据
func withValue(arr [NumOfElems]Content) int {
	//fmt.Println(&arr[2])
	return 0
}

// 传指针
func withReference(arr *[NumOfElems]Content) int {
	//b := *arr
	//fmt.Println(&arr[2])
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumOfElems]Content
	//fmt.Println(&arr[2])
	withValue(arr)
	withReference(&arr)
}

func BenchmarkPassingArrayWithValue(b *testing.B) {
	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

func BenchmarkPassingArrayWithRef(b *testing.B) {
	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
}

/*
性能差异：
BenchmarkPassingArrayWithValue-8  76          17875772 ns/op
BenchmarkPassingArrayWithRef-8    1000000000   0.2964 ns/op

*/

/*
打开GC日志
只要在程序执行前加上环境变量GODEBUG=gctrace=1
如 GODEBUG=gctrace=1 go test -bench=.  不知道为啥运行不了
   GODEBUG=gctrace=1 go run main.go
日志详细信息参考：https://godoc.org/runtime

*/

/*


go test -bench=BenchmarkPassingArrayWithRef -trace=trace_ref.out
go test -bench=BenchmarkPassingArrayWithValue -trace=trace_val.out
生成两个.out文件
go tool trace trace_ref.out打开浏览器
查看goroutines、heap、threads、GC等信息

*/
