package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

//添加错误机制
//给errors.New()声明一个变量应该使用err或Err开头
var errLessThanTwo = errors.New("n should be not less than 2")
var errLargeThanHundred = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, errLessThanTwo
	}
	if n > 100 {
		return nil, errLargeThanHundred
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-5); err != nil {
		if err == errLessThanTwo {
			fmt.Println("It is less.")
		}
		if err == errLargeThanHundred {
			fmt.Println("Tt is large.")
		}
		//t.Error(err)
	} else {
		t.Log(v)
	}
	//t.Log(GetFibonacci(10))
}

func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

//推荐第二个，及早失败，避免嵌套
func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}
