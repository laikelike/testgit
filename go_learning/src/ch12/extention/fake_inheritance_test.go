package extention

import (
	"fmt"
	"testing"
)

//假继承，不能重载父类的方法
type Pet1 struct {
}

func (p *Pet) Speak1() {
	fmt.Print("...")
}
func (p *Pet) SpeakTo1(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog1 struct {
	Pet //匿名嵌套类型
}

//重载不了
func (d *Dog1) Speak1() {
	fmt.Print("wang wang")
}
func TestDog1(t *testing.T) {
	dog := new(Dog)
	//dog.Speak()
	dog.SpeakTo("laike") //如果重载成功打印出的应是wang，而不是...
}
