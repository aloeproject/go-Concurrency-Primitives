package cond

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func condErr1() {
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
	/*
	调用wait不加锁
	cond.Wait实现是，当前调用者进入notify队列后，会释放锁
	其他调用者会去争抢这把锁，如果调用者调用之前不加锁，可能unlock一个未加锁的Locker
	 */
	//c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Printf("裁判员被唤醒一次")
	}
	//c.L.Unlock()

	log.Println("所有运动员准备就绪，比赛开始")
}



func condErr2() {
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
	/*
	waiter 被唤醒不等于条件满足，他只是被唤醒而言，条件有可能满足，有可能不满足
	我们需要进一步的检查。
	 */
	//for ready != 10 {
		c.Wait()
		log.Printf("裁判员被唤醒一次")
	//}
	c.L.Unlock()

	log.Println("所有运动员准备就绪，比赛开始")
}