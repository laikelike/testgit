/*函数是一等公民
1.函数可以有多个返回值
2.所有参数都是传值：slice,map,channel因为是共享空间会有传引用的错觉。
3.函数可以作为变量的值，例如map的值
4.函数可以作为参数和返回值
*/
package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//函数多返回值
//第一个小括号里是传入参数类型，第二个是返回参数类型
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

//专门用来计时的函数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n) //调用内部函数
		//输出内部函数运行的时间
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

//可变长参数,用在参数定义时的"..."表明此参数可变长
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

//延迟执行函数类似try finally函数
func Clear() {
	fmt.Println("Clear resources.")
}
func TestDefer(t *testing.T) {
	//defer后的代码无论如何都会执行，即使有错误panic，可以用来释放锁
	defer Clear() //函数返回前才执行，类似于finally{}
	fmt.Println("start")
	//panic("error in programing")
	//fmt.Println("end") //不可达代码
}
