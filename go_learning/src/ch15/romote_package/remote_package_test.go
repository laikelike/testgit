package remote_test

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

/*
引入github中的外部包在自己项目的文件目录下，
cd到ch15下的main包中用go get -u github.com/easierway/concurrent_map
然后就能运行了

Go Modules现在是官方推荐的依赖管理工具
1.出现的目的之一是解决GOPATH问题，以前下古墓必须在$GOPATH/src里进行，
现在GO允许在其之外任何目录下使用go.mod创建项目
2.目前所有模块版本数据都缓存在$GOPATH/pkg/mod和$GOPATH/pkg/sum下
3.可以用go clean -modcache清理所有已缓存的木块版本数据
4.可用go env -w GOPROXY=https://goproxy.cn,direct修改goproxy
目前用的是https://goproxy.io好像也还行？先不动，以后有需要再改
5.创建一个有module依赖管理的项目步骤创建项目后
在有代码的目录下(一般是项目的根目录下,如go_leanrning)go mod init src 生成了go.mod文件
在src下执行go build 生成go.sum
当在需要调用git上包时，在代码目录下执行go get -u github:..命令
将此包下载到了默认GOPATH下的pkg下的mod下的github.com下
此时go.mod中就会有一条github命令

还要了解go replace
*/
