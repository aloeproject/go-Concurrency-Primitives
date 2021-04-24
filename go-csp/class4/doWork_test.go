package class4

import (
	"fmt"
	"testing"
	"time"
)

func TestDoWork(t *testing.T) {
	doWork := func(ch <-chan int) <-chan interface{} {
		result := make(chan interface{})
		go func() {
			defer fmt.Println("work done")
			defer close(result)
			for i := range ch {
				result <- i
			}
		}()

		return result
	}

	//如果往doWork传nil，会产生阻塞
	//导致goroutine泄露
	ch := doWork(nil)
	<-ch
}

func TestDoWork2(t *testing.T) {
	doWork := func(done chan struct{}, ch <-chan int) <-chan interface{} {
		result := make(chan interface{})
		go func() {
			defer fmt.Println("work done")
			defer close(result)
			for {
				select {
				case v := <-ch:
					result <- v
					//增加done，外面控制它是否阻塞
				case <-done:
					return
				}
			}
		}()
		return result
	}

	done := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	ch := doWork(done, nil)
	<-ch
	fmt.Println("finish")
}
