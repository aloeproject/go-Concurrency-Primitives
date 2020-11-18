package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestCMutex_Lock(t *testing.T) {
	lock := NewCMutex()
	var ct int
	for i := 0; i < 100000; i++ {
		go func() {
			lock.Lock()
			ct++
			lock.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Print(ct)
}

func TestCMutex_TryLock(t *testing.T) {
	lock := NewCMutex()
	go func() {
		lock.Lock()
		lock.Unlock()
	}()
	time.Sleep(1 * time.Millisecond)
	go func() {
		fmt.Println(lock.TryLock())
	}()
	time.Sleep(1 * time.Second)
}