package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

//对数据的封装
type Employee struct {
	Id   string
	Name string
	Age  int
}

/*对行为的封装*/

//在实例对应方法被调用时,实例的成员会对值复制，不推荐使用
// func (e Employee) String() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

// 使用取地址符，共用同一块内存空间，地址是相同的,避免了内存拷贝推荐这种方法
func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) //返回指针
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T) {
	//e := Employee{"0", "Bob", 20}
	e := &Employee{"0", "Bob", 20} //用指向实例的指针也可以调用
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
