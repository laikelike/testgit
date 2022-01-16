package extention

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Println("wang wang!")
}
func (d *Dog) SpeakTo(host string) {

	//会调用Pet中的speakto中的p.speak()，所以输出是...
	d.p.SpeakTo(host)

	//必须自己实现speak方法，此为复合
	// d.Speak()
	// fmt.Println("", host)
}
func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("laike")
}
