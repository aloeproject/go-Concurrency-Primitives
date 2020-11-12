package trylock

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	//这些状态是 sync包里面的状态
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type Mutex struct {
	sync.Mutex
}

/*
判断锁是否被锁，如果锁返回true
如果没有上锁，则获取这个锁返回获取状态
*/
func (m *Mutex) TryLock() bool {
	//如果能抢到锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	new1 := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new1)
}

func (m *Mutex) Count() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift //得到等待者数量
	v = v + (v + mutexLocked) //再加上锁持有者的数量，0或者1
	return int(v)
}

//当前是否被锁
func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	//为什么不直接判断 state == mutexLocked
	return state&mutexLocked == mutexLocked
}

//是否有等待者被唤醒
func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

//锁释放处于饥饿状态
func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}
