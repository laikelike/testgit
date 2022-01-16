package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	//if语句支持变量赋值
	/*  if var declaration; conditon  前面是赋值，后面直接用赋值的结果判断*/
	// b := 1
	// if a := 1 == b; a {
	// 	t.Log("1==1")
	// } else {
	// 	t.Log("1==1")
	// }

	//v是方法本身的返回值，err是错误
	// if v, err := someFun(); err == nil {
	// 	t.Log("abc") 若没有错误执行abc
	// }else{		若有错误，执行cde
	// 	t.Log("cde")
	// }
}

// GO中的swich语句不需要case
func TestSwitchMultiCase(t *testing.T) {
	for i := 1; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseConditon(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}
}
