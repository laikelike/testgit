package flexible_reflect_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

//reflect.DeepEqual()可以用来比较两个map或切片
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//fmt.Println(a == b) //map之间不能比较，只能和空比较
	fmt.Println(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 1, 3}
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	// c1 := Customer{"1", "Mike", 40}
	// c2 := Customer{"1", "Mike", 40}

}

/*
Customer和Employee可以利用反射机制用一个方法给这两个结构体都可以
赋值
*/

func fillBySettings(st interface{}, settings map[string]interface{}) error {

	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		//Element() 获取指针指向的值
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}
	if settings == nil {
		return errors.New("setting is nil.")
	}
	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 40}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

/*
反射
提高了程序的灵活性
降低了程序的可读性
降低了程序的性能

*/
