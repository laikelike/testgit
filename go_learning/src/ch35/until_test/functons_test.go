package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("\ninput is %d, the expected is %d, the actual is %d",
				inputs[i], expected[i], ret)
		}
	}
}

//Fail,Error:该测试失败，剩下测试继续，其他测试继续执行
func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

//FailNow,Fatal:该测试失败，剩下测试中止，其他测试继续执行
func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("End")
}

//下载断言包，使用断言比较测试
func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}
