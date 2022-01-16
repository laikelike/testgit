package client

import (
	"WorkSpace/go_learning/src/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
	t.Log(series.Sequare(5))
}

/*调用自定义函数
1.如果项目不在GOPATH环境变量的路径里，需要添加环境变量
2.包文件的函数名第一个字母要大写，这样外部可以调用，同一个package无所谓
3.包的.go文件必须放在一个独立的文件夹下(如serices)
4.对serices文件使用go build 和go install 命令
5.检查在go_learning目录下生成pkg文件，文件中有series.a文件
问题：课程中import "ch15/series" 意思是说添加GOPATH环境变量后src属于
GoPath的一部分不用写出，所以引入的是 "ch15/series",但是我引入的是
import "WorkSpace/go_learning/src/ch15/series"并且发现无论是否添加
go_learning环境变量都可以编译成功，为啥？
自己猜测问题出在一开始构建工作区时，构建在了workSpace文件夹下，但解释不通
为啥不添加环境变量也可以编译通过？
*/
