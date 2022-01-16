package series

import "fmt"

/*
1.在mian执行前，所有依赖的package的init方法都会被执行
2.不同包的init函数按照包导入的依赖关系决定执行顺序
3.每个包可以有多个init函数
4.包的每个元文件也可以有多个init函数，这点比较特殊
*/
func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}

//文件名首字母大写，包外函数可以调用类似于Java中的公有函数
//首字母小写，类似于私有函数
func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

func Sequare(n int) int {
	return n * n
}
