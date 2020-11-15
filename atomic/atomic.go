package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Add() {
	var i32 int32
	i32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&i32, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(i32)
}

func Load() {
	var i32 int32
	i32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&i32, 1)
			fmt.Println(atomic.LoadInt32(&i32))
			wg.Done()
		}()
	}
	wg.Wait()
}

func CompareAndSwap() {
	var i32 int32
	i32 = 0
	//true
	fmt.Println(atomic.CompareAndSwapInt32(&i32, 0, 1))
	fmt.Println(i32)
	//false
	fmt.Println(atomic.CompareAndSwapInt32(&i32, 0, 1))
}

func store()  {
	var i32 int32
	i32 = 0
	atomic.StoreInt32(&i32,1)
	fmt.Println(i32)
}
