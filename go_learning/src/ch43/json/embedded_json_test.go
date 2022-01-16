package jsontest

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
		"basic_info":{
		"name":"Mike",
		"age":30
		},
		"job_info":{
		"skills":["Java","Go","C"]
		}
	}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	//用字符串对空对象进行填充
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}

/*
内置的Json解析
利用反射实现，通过FeildTag来标识对应的json值
性能不太行
*/
