package ch51

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const numbers = 100

// 通过fmt的Sprintf连接字符串
func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

// go 1.10之后更新的strings.Builder,使用方法和buffer类似，推荐！
func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var builder strings.Builder
		for i := 0; i < numbers; i++ {
			builder.WriteString(strconv.Itoa(i))
		}
		_ = builder.String()
	}
	b.StopTimer()
}

// 通过bytes.Buffer连接字符串，利用buffer的连续存储空间
func BenchmarkBytesBuf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var buf bytes.Buffer
		for i := 0; i < numbers; i++ {
			buf.WriteString(strconv.Itoa(i))
		}
		_ = buf.String()
	}
	b.StopTimer()
}

// 通过Stringadd连接
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s += strconv.Itoa(i)
		}
	}
	b.StopTimer()
}

/*
高效字符串拼接


BenchmarkSprintf-8                 74961             23679 ns/op
BenchmarkStringBuilder-8         1000000              1679 ns/op
BenchmarkBytesBuf-8               872919              1196 ns/op
BenchmarkStringAdd-8              183178              6597 ns/op

所以推荐StringBuilder
*/
