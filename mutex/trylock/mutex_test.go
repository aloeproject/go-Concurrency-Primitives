package trylock

import (
	"fmt"
	"testing"
	"time"
)

func TestMutex_TryLock(t *testing.T) {
	/*
		先加锁，如果获取锁状态失败 那么打印 get lock fail
	*/
	lo := new(Mutex)
	go func() {
		lo.Lock()
		time.Sleep(2 * time.Second)
		lo.Unlock()
	}()
	time.Sleep(1 * time.Microsecond)

	pass := lo.TryLock()
	if pass {
		fmt.Println("get lock ok")
		lo.Unlock()
		return
	}
	fmt.Println("get lock fail")
}

func Test_Print(t *testing.T) {
	var mu Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("waitings:%d,isLocked:%t,woken:%t,starving:%t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
