package benchmark

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("12345", ret)
}

func TestConcatStringBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {

	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StartTimer()
}

func BenchmarkConcatStringBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	var buf bytes.Buffer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StartTimer()
}

/*
1.基准测试的代码文件必须以_test.go结尾
2.基准测试的函数必须以Benchmark开头
3.基准测试函数必须接受一个指向Benchmark类型的指针作为唯一参数，即比如func BenchmarkMapkeys1(b *testing.B)
4.基准测试函数不能有返回值
5.b.ResetTimer是重置计时器，这样可以避免for循环之前的初始化代码的干扰
6.最后的for循环很重要，被测试的代码要放到循环里
7.b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能
8.不要修改b.N，因为基准测试是很智能的，它会自动根据你要测试函数的运行时间动态调整该值
9.b.N 从 1 开始，如果基准测试函数在1秒内就完成 (默认值)，则 b.N 增加，并再次运行基准测试函数。b.N 在近似这样的序列中不断增加；1, 2, 3, 5, 10, 20, 30, 50, 100,1000,20000 等等。 基准框架试图变得聪明，如果它看到当b.N较小而且测试很快就完成的时候，它将让序列增加地更快



在cmd命令行中运行
go test -bench=.
"."代表所有程序，指定某一方法的化后面加方法名

方法名(-8:使用线程数)                运行次数  每次操作所耗时间
BenchmarkConcatStringByAdd-8        8742608  154.2 ns/op
BenchmarkConcatStringBytesBuffer-8  27331426  8.58 ns/op

go test -bench=. -benchmem
1.参数-bench，它指明要测试的函数；点字符意思是测试当前所有以Benchmark为前缀函数
2.参数-benchmem，性能测试的时候显示测试函数的内存分配大小，内存分配次数的统计信息
3.参数-count n,运行测试和性能多少此，默认一次

方法名                             运行次数   每次操作所耗时间   每次操作申请多少byte内存  每次操作申请几次内存
BenchmarkConcatStringByAdd-8       7488182   141.1 ns/op      16 B/op                  4 allocs/op
BenchmarkConcatStringBytesBuffer-8 34843508  35.26 ns/op      15 B/op                  0 allocs/op

*/
