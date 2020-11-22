package gerrgroup

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

/*
出现多个错误只打印第一个错误
*/
func errGroup1() {
	var g errgroup.Group

	g.Go(func() error {
		time.Sleep(3 * time.Second)
		str := "task1"
		fmt.Println(str)
		return nil
	})

	g.Go(func() error {
		time.Sleep(3 * time.Second)
		str := "task2"
		fmt.Println(str)
		return fmt.Errorf("error %s", str)
	})

	g.Go(func() error {
		time.Sleep(3 * time.Second)
		str := "task3"
		fmt.Println(str)
		return fmt.Errorf("error %s", str)
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("finish")
}

func pipeline() {
	var g errgroup.Group
	inputNum := make(chan int)
	g.Go(func() error {
		defer close(inputNum)
		for i := 0; i < 10000; i++ {
			inputNum <- 1
		}
		return fmt.Errorf("intput error")
	})

	resNum := make(chan int)
	lock := sync.Mutex{}
	var sum int
	//启动多个goroutine
	for i := 0; i < 5000; i++ {
		g.Go(func() error {
			for v := range inputNum {
				//加锁防止 保证加锁原子性
				lock.Lock()
				sum += v
				lock.Unlock()
				time.Sleep(1 * time.Second)
				resNum <- sum
			}
			return nil
		})
	}

	go func() {
		g.Wait()
		close(resNum)
	}()

	var result int
	for v := range resNum {
		result = v
	}

	fmt.Println(result)
	//再次调用 wait还能捕获到错误
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

func cancelWork() {
	/*
		task2发生任务错误，任务全部结束
	*/
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		name := "task1"
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, " cancel")
				return nil
			default:
				fmt.Println(name, "do some things")
				time.Sleep(3 * time.Second)
			}
		}
	})

	g.Go(func() error {
		name := "task2"
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, " cancel")
				return nil
			default:
				fmt.Println(name, "do some things")
				time.Sleep(1 * time.Second)
				//task2 发生了错误
				return fmt.Errorf("%s error", name)
			}
		}

	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("finish")
}
