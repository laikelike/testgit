package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f) //传入地址不匹配
}

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

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	//按名字获取成员
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format")) //键值对形式
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age: ", e)
}

/*
reflect.TypeOf返回类型(reflect.Type)
reflect.ValueOf返回值(reflect.Value)
可以从reflect.Value获得类型

判断类型-kind()，返回instance的类型
按照名字访问结构的成员：
reflect.ValueOf(*e).FiledByName("Name")
按名字访问结构的方法
reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})

Struct Tag
type BaseicInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`//单括号引起来的是struct Tag
}
访问StructTag两种方法
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok{
		t.Error("Failed to get 'Name' field.")
	}else{
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}

Reflect.Type和Reflect.Value都有FieldByName方法，但有区别
*/
