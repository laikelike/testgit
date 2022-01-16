package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1]) //用来传入命令行参数
	}
	fmt.Println("hello")
	fmt.Println("hello" + "laike")
	fmt.Println("--------------------------")
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "code = %d & endDate = %s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	os.Exit(-1) //返回值

}
