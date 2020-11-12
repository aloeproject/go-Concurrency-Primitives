package trylock

import (
	"fmt"
	"testing"
	"time"
)

func TestMutex2_TryLock(t *testing.T) {
	/*
		先加锁，如果获取锁状态失败 那么打印 get lock fail
	*/
	lo := new(Mutex2)
	go func() {
		lo.Lock()
		time.Sleep(2 * time.Second)
		lo.Unlock()
	}()
	time.Sleep(1 * time.Millisecond)

	pass := lo.TryLock(1 * time.Second)
	if pass {
		fmt.Println("get lock ok")
		lo.Unlock()
		return
	}
	fmt.Println("get lock fail")
}