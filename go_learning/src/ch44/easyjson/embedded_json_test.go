package easyjson

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

func TestEasyJson(t *testing.T) {
	e := Employee{}
	e.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(e)
	if v, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(v))
	}
}

func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err = json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := Employee{}
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			b.Error(err)
		}
		if _, err = json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}

/*
go test -bench=. -benchmem

BenchmarkEmbeddedJson-8     439711  3879 ns/op    568 B/op    10 allocs/op
BenchmarkEasyJson-8         707023  2093 ns/op    380 B/op    7 allocs/op
*/
