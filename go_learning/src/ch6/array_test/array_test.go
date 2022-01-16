package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int //默认初始化为0
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5} //若懒得数有几个元素用三个点代替
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}
func TestMoreDimensionArrayInit(t *testing.T) {
	b := [][]int{} //切片
	c := [2][2]int{{1, 2}, {3, 4}}
	e := []int{4, 5}
	d := []int{6, 7}
	b = append(b, e)
	b = append(b, d)
	b = append(b, d)
	t.Log(c)
	t.Log(b)
	val := b[1][1] //数组下标从0开始
	t.Log(val)
}
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, e := range arr3 { //idx是数组下标，e是元素值
		t.Log(idx, e)
	}
	t.Log("------------------------------")
	for _, e := range arr3 { //不需要idx时,用下划线代替
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5, 6}
	arr3_sec := arr3[3:] //数组下标3到结尾
	t.Log(arr3_sec)
	arr3_sec1 := arr3[:]
	t.Log(arr3_sec1)
	arr4_sec1 := arr3[3:5] //数组下标3到4，不包括5
	t.Log(arr4_sec1)
}
