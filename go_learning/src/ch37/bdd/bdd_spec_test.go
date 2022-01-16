package bdd_test

import (
	"testing"
	//import包路径前加一个点，证明这个包里的所有方法在当前名字空间
	//不需要引用可直接使用
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	// only pass t into top-level Convey calls
	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4
		Convey("When add the two numbers", func() {
			c := a + b

			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0) //断言
			})
		})
	})
}
