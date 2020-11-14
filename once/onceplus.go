package once

import (
	"sync"
	"sync/atomic"
)

type OncePlus struct {
	done int32
	m    sync.Mutex
}

func (o *OncePlus) Do(f func() error) error {
	if atomic.LoadInt32(&o.done) == 0 {
		return o.doSlow(f)
	}
	return nil
}

/**
如果 初始返回错误，再次调用可以被初始化
 */
func (o *OncePlus) doSlow(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreInt32(&o.done, 1)
		}
	}
	return err

}
