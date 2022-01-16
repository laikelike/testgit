package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len ma = %d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2 = %d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3 = %d", len(m3)) //map不通用cap()求容量
}

//在访问的key不存在时，仍会返回零值，不能通过返回nil来判断元素是否存在
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	//第一个返回参数是值，第二个是是否为空，都是布尔值
	if value, isnull := m1[3]; isnull {
		t.Logf("key 3's value is %d", value)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for key, value := range m1 {
		t.Log(key, value)
	}
}
