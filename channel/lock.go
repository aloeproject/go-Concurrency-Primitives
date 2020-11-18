package channel

import "time"

/*
使用channel 实现一个互斥锁
容量为1 的chan ，谁抢到谁不被阻塞
*/

type CMutex struct {
	ch chan struct{}
}

func NewCMutex() *CMutex {
	return &CMutex{ch: make(chan struct{}, 1)}
}

func (c *CMutex) Lock() {
	c.ch <- struct{}{}
}

func (c *CMutex) Unlock() {
	select {
	case <-c.ch:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (c *CMutex) TryLock() bool {
	select {
	case c.ch<- struct{}{}:
		return true
	default:
		return false
	}
}

func (c *CMutex) LockTimeout(d time.Duration) bool {
	t := time.NewTicker(d)
	select {
	case  c.ch<- struct{}{}:
		t.Stop()
		return true
	case <-t.C:
		return false
	}
}

func (c *CMutex) IsLocked() bool {
	return len(c.ch) == 1
}
