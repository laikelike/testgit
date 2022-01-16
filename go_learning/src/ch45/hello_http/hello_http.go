package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Go world!")
	})
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})
	//启动后ctrl+c退出

}

/*
路由规则
1.URL分为两种，末尾时/:表示一个子树，后面可以跟其他子路径，末尾不是
/，表示一个叶子，时固定的路径。(以/结尾的URL可以匹配它的任何子路径，
比如/images会匹配/images/cute-cat.jpg)
2.它采用最长匹配原则，如果有多个匹配，一定采用匹配路径最长的哪个进行处理
3.如果没有找到任何匹配项，会返回404错误


Defalut Router
http.ListenAndServe(":8080", nil)
	func (sh serverHandler) ServeHttp(rw RespnseWriter, req *Request)  {
		handler := sh.srv.Handler
		if handler == nil {
			handler = DefaultServeMux//使用缺省的Router
		}
		if req.requestURI == "*" && req.Method == "OPTIONS"{
			handler = globalOptionsHandler{}
		}
		handler.ServeHTTP(rw, req)

	}


*/
