package interface_test

import "testing"

//接口是非入侵性的，实现不依赖于接口定义
//接口的定义可以包含在接口使用者包内

//定义接口
type Programmer interface {
	WriteHelloWorld() string
}

//定义接口的实现
type GoProgrammer struct {
}

//duck type方法签名看起来是一样的
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Helllo World\")"
}

func TestClient(t *testing.T) {
	//var p Programmer
	var p Programmer = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
