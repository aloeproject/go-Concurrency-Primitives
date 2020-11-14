package cond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLimitedQueue_Push(t *testing.T) {
	obj := NewLimitedQueue(3)
	obj.Push(1)
	obj.Push(2)
	//没有消费会被死锁
	obj.Push(3)
}

func TestLimitedQueue_Push2(t *testing.T) {
	obj := NewLimitedQueue(3)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	//生产者
	go func() {
		for i := 0; i < 5; i++ {
			obj.Push(i + 1)
		}
		wg.Done()
	}()

	//消费者
	go func() {
		for {
			val, err := obj.Pop()
			if err != nil {
				t.Log(err)
				break
			}
			fmt.Println(val)
		}
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 3)
		//关闭队列
		obj.Close()
	}()
	wg.Wait()
}
