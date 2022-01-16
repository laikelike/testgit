package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func GetFibonacciSerie(n int) []int {
	ret := make([]int, 2, n)
	ret[0] = 1
	ret[1] = 1
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func createFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int
	for i := 0; i < 1000000; i++ {
		fbs = GetFibonacciSerie(50)
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs)))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/fb", createFBS)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/*
通过HTTP方式输出Profile
1.简单，适用于持续性运行的应用
2.在应用中导入import_"net/http/pprof",并启动http server即可
3.http://<host>:<post>/debug/pprof/
4.go tool pprof _http://<host>:<post>/debug/pprof/profile?seconds=10(默认值为30秒)
5.go-torch-seconds 10 http://<host>:<post>/debug/pprof/profile


例如：go tool pprof http://127.0.0.1:8081/debug/pprof/profile
这里会有一个默认30秒的采样时间，在采样时间内再次请求网页，查看资源占用情况(我请求了3次)

使用命令“top”
 flat  flat%   sum%        cum   cum%
240ms 17.52% 17.52%      240ms 17.52%  runtime.stdcall1
190ms 13.87% 31.39%      190ms 13.87%  runtime.stdcall3
80ms  5.84% 37.23%       80ms  5.84%  runtime.stdcall2
60ms  4.38% 41.61%      190ms 13.87%  main.GetFibonacciSerie
60ms  4.38% 45.99%       80ms  5.84%  runtime.scanobject

使用命令“list GetFibonacciSerie”

60ms      190ms (flat, cum) 13.87% of Total
         .          .      6:   "net/http"
         .          .      7:   _ "net/http/pprof"
         .          .      8:)
         .          .      9:
         .          .     10:func GetFibonacciSerie(n int) []int {
         .      130ms     11:   ret := make([]int, 2, n)
         .          .     12:   ret[0] = 1
         .          .     13:   ret[1] = 1
      30ms       30ms     14:   for i := 2; i < n; i++ {
      30ms       30ms     15:           ret = append(ret, ret[i-2]+ret[i-1])
         .          .     16:   }
         .          .     17:   return ret
         .          .     18:}
		 .          .     19:
         .          .     20:func index(w http.ResponseWriter, r *http.Request) {
(pprof) list -cum
*/
