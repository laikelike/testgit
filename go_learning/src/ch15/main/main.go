package main

import (
	"fmt"
	cm "github.com/easierway/concurrent_map"
)

func main() {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	fmt.Println(m.Get(cm.StrKey("key")))
}
