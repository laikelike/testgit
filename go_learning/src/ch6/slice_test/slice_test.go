package slice_test

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int            //中括号里没有...就是切片
	t.Log(len(s0), cap(s0)) //len是已使用大小，cap是连续空间容量
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5) //创建某一类型的变量(类型,length,capacity)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
}
func TestSliceGrowing(t *testing.T) {
	s := []int{} // 切片二倍自增空间
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unkown"
	t.Log(year)
}
func TestSliceComparing(t *testing.T) {
	//切片不可以比较
	// a:= []int{1,2,3,4}
	// b := []int{1,2,3,4}
	// if a==b {
	// 	t.Log("equal")
	// }
}
