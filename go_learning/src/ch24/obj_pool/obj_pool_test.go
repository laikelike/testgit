package objpool_test

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj //用于缓冲可重用对象
}

func NewObjPool(numOfObj int) *ObjPool {
	ObjPool := ObjPool{}
	ObjPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		ObjPool.bufChan <- &ReusableObj{}
	}
	return &ObjPool
}

//若需要放置不同类型的对象可以将*ReusableObj类型改成空接口
//interface{}类型，若是此类型得到对象时需要断言其类型
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	//slow response(慢反应)比quick failure(快失败)更糟糕
	case <-time.After(timeout): //超时控制，
		return nil, errors.New("\ntime out")
	}
}

//得到对象的方法和取释放对象的方法必须是池的方法加上(p *ObjPool)
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow") //溢出
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	//尝试放置超过线程池
	if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
		t.Error(err)
	}
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
}

/*
buffer channel实现对象池
应用时需要思考：当需要新的对象时，使用对象池(有锁的开销)或是新建
一个对象哪个开销大，有的时候不见得使用对象池比new对象更快

*/
