package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const (
	row = 10000
	col = 10000
)

func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(100000)
		}
	}
}

func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}

func main() {
	//创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	// 获取系统信息
	if err := pprof.StartCPUProfile(f); err != nil { //监控CPU
		log.Fatal("could nor start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	//主逻辑区，进行一些简单的代码运算
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	runtime.GC()                                       // GC, 获取最新的数据信息
	if err := pprof.WriteHeapProfile(f1); err != nil { //写入内存信息
		log.Fatal("could not write memory profile: ", err)
	}
	f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create groutine profile: ", err)
	}
	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		log.Fatal("could not write groutine profile: ")
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()
}

/*
通过文件方式输出Profile
1.灵活性高，适用于特定代码段的分析
2.通过手动调用runtime/pprof的API
3.API相关文档https://studygolang.com/static/pkgdoc/pkg/runtime_pprof.htm
4.GO支持多种Profile https://golang.org/src/runtime/pprof/pprof.go
*/

/*
pprof包有两个："net/http/pprof"(系统默认导入这个)和"runtime/pprof"
net/http/pprof中只是使用runtime/pprof包来进行封装了一下，并在http端口上暴露出来
涉及web服务时用net那个，只是应用程序用runtime那个
*/

/*
运行使用系统内置性能分析工具：pprof
命令行运行go build prof.go 命令生成了cpu.prof、gortine.prof和
mem.prof三个文件
运行 go tool pprof 编译出的二进制文件名 想查看的prof文件 例如：go tool pprof prof cpu.prof 进入一个交互式控制台
“exit”命令退出
用“top”命令查看从大到小CPU使用情况

函数执行时间 所占比例				(cumulation累积)
flat        flat%  		 sum%        cum  		 cum%
420ms     42.00% 		42.00%      600ms 		60.00%  math/rand.(*Rand).Int31n
220ms     22.00% 		64.00%      900ms 		90.00%  main.fillMatrix
80ms      8.00% 		72.00%      680ms 		68.00%  math/rand.(*Rand).Intn
80ms      8.00% 		80.00%      160ms 		16.00%  math/rand.(*rngSource).Int63
80ms      8.00%		 	88.00%       80ms  		8.00%  math/rand.(*rngSource).Uint64 (inline)
20ms      2.00% 		90.00%       20ms  		2.00%  main.calculate

cum和cum%代表这个函数本身和所调用的函数总体所占使时间和比例

用“list 函数名”(可自动进行最大匹配，不一定需要敲全)查看详细函数资源占用情况 例如： list fillMatrix

 		220ms      900ms (flat, cum) 90.00% of Total
         .          .     14:)
         .          .     15:
         .          .     16:func fillMatrix(m *[row][col]int) {
         .          .     17:   s := rand.New(rand.NewSource(time.Now().UnixNano()))
         .          .     18:   for i := 0; i < row; i++ {
      50ms       50ms     19:           for j := 0; j < col; j++ {
     170ms      850ms     20:                   m[i][j] = s.Intn(100000)
         .          .     21:           }
         .          .     22:   }
         .          .     23:}
         .          .     24:
         .          .     25:func calculate(m *[row][col]int) {

用命令“svg”生成svg图，用浏览器打开查看

go-torch cpu.prof  因为代码是针对32位系统的所以没生成火炬图
 %1 is not a valid Win32 application.
*/
