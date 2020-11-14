package cond

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func broadcast() {
	c := sync.NewCond(new(sync.Mutex))
	var ready int
	//运动员数量
	playerNum := 10
	for i := 0; i < playerNum; i++ {
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

	wg := new(sync.WaitGroup)
	wg.Add(2)
	/*
		两个裁判员同时监听,当调用广播时都被唤醒
		会打印

		运动员0,准备就绪
	 	裁判员1被唤醒一次
	 	裁判员2被唤醒一次
		所有运动员准备就绪，比赛开始
	*/
	go func() {
		defer wg.Done()
		c.L.Lock()
		for ready != playerNum {
			c.Wait()
			log.Printf("裁判员1被唤醒一次")
		}
		c.L.Unlock()
	}()

	go func() {
		defer wg.Done()
		c.L.Lock()
		for ready != playerNum {
			c.Wait()
			log.Printf("裁判员2被唤醒一次")
		}
		c.L.Unlock()
	}()
	wg.Wait()

	log.Println("所有运动员准备就绪，比赛开始")
}
