package fanIn

import (
	"sync"
	"time"
)

func work(done chan struct{}, value ...int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for _, v := range value {
			select {
			case <-done:
				return
			case result <- v:
				//这里模拟随机的工作耗时
				//n := rand.Intn(2)
				//time.Sleep(time.Duration(n) * time.Second)
				time.Sleep(2 * time.Second)
			}
		}
	}()
	return result
}

func work2(done chan struct{}, value ...int) <-chan int {
	result := make(chan int)
	close(result)
	return result
}

func fanIn(done chan struct{}, ch ...<-chan int) <-chan int {
	result := make(chan int)
	var wg sync.WaitGroup
	mul := func(val <-chan int) {
		defer wg.Done()
		for v := range val {
			select {
			case <-done:
				return
			case result <- v:
			}
		}
	}

	for _, v := range ch {
		wg.Add(1)
		go mul(v)
	}

	go func() {
		//等待全部写入完成，然后进行关闭
		wg.Wait()
		close(result)
	}()

	return result
}

func fanOr(done chan struct{}, ch ...<-chan int) <-chan int {
	result := make(chan int)
	var wg sync.WaitGroup
	mul := func(val <-chan int) {
		defer wg.Done()
		//这里也可以加上是否为空判断
		if val == nil {
			return
		}
		for {
			select {
			case <-done:
				return
				/*
					可能你读取的 channel 已经关闭，需要再加判断
					但是这里不需要过早优化
				*/
			case v1, ok := <-val:
				if ok == false {
					return
				}

				select {
				case result <- v1:
				case <-done: //不太懂这一句
				}

			}
		}
	}

	for _, v := range ch {
		wg.Add(1)
		go mul(v)
	}

	go func() {
		//等待全部写入完成，然后进行关闭
		wg.Wait()
		close(result)
	}()

	return result
}
