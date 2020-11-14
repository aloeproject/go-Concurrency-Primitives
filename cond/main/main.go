package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
type Cond

func(c Cond)Broadcast()
，允许调用者 Caller 唤醒一个等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，
显然无需通知 waiter；如果 Cond 等待队列中有一个或者多个等待的 goroutine，
则需要从等待队列中移除第一个 goroutine 并把它唤醒。在其他编程语言中，比如 Java 语言中，
Signal 方法也被叫做 notify 方法。

func(c Cond)Signal()
，允许调用者 Caller 唤醒所有等待此 Cond 的 goroutine。
如果此时没有等待的 goroutine，显然无需通知 waiter；
如果 Cond 等待队列中有一个或者多个等待的 goroutine，
则清空所有等待的 goroutine，并全部唤醒。在其他编程语言中，
比如 Java 语言中，Broadcast 方法也被叫做 notifyAll 方法。

func(c Cond) Wait(),会把调用者 Caller 放入 Cond 的等待队列中并阻塞，
直到被 Signal 或者 Broadcast 的方法从等待队列中移除并唤醒。
*/

func cond() {
	c := sync.NewCond(new(sync.Mutex))
	var ready int
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			//加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员%d,准备就绪\n", i)
			//广播通知所有的等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Printf("裁判员被唤醒一次")
	}
	c.L.Unlock()

	log.Println("所有运动员准备就绪，比赛开始")
}


/*
channel 方式实现
 */
func channel() {
	ch := make(chan bool)
	lock := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	var ready int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			lock.Lock()
			ready++
			ch <- true
			if ready == 10 {
				close(ch)
			}
			lock.Unlock()
			log.Printf("运动员%d,准备就绪\n", i)
		}(i)
	}

	go func() {
		for range ch {
			log.Printf("裁判员被唤醒一次")
		}
	}()
	wg.Wait()
	log.Println("所有运动员准备就绪，比赛开始")
}

func main() {
	cond()
	//channel()
}
