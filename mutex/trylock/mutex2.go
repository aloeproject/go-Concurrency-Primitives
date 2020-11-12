package trylock

import (
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type Mutex2 struct {
	sync.Mutex
}

/**
添加超时机制，如果在时间内没有获取到锁则返回false
*/
func (m *Mutex2) TryLock(timeout time.Duration) bool {
	//如果能抢到锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	tk := time.NewTicker(timeout)

	/*
	目前还不知道这单代码具体，不可靠原因是在哪里
	 */
	for {
		select {
		case <-tk.C:
			return false
		default:
			old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
			if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
				time.Sleep(1 * time.Microsecond)
			} else {
				new1 := old | mutexLocked
				return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new1)
			}
		}
	}
}
