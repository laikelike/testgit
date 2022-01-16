package emptyinterface

import (
	"fmt"
	"testing"
)

func Dosomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("This type is Integer", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("This type is string", s)
	// 	return
	// }
	// fmt.Println("Unknow Type")

	//switch用在这里更清晰明了
	switch v := p.(type) {
	case int:
		fmt.Println("This type is Integer", v)
	case string:
		fmt.Println("This type is string", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	Dosomething(10)
	Dosomething("10")
}

/*Go接口最佳实践
1.倾向于使用小的接口定义，很多接口只包含一个方法
type Reader interface{
	Read(p []byte)(n int,err error)
}
type Writer interface{
	Write(p []byte)(n int,err error)
}
2.较大的接口定义，可以由多个小接口定义组合而成
type ReadWriter interface{
	Writer
	Reader
}
3.只依赖于必要功能的最小接口
func StoreData(reader Reader)error{
	.......
}
*/
