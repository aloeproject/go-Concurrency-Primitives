package context1

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func withValue() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key1", "001")
	ctx = context.WithValue(ctx, "key2", "002")
	ctx = context.WithValue(ctx, "key3", "003")
	/*
		查找方式  从子context往上查找
			key3->key2->key1
	*/
	fmt.Println(ctx.Value("key1"))
}

func withCancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("work finish")
				return
			default:
				fmt.Println("work run")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func withCancel2() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	for i := 0; i < 5; i++ {
		go func(ctx context.Context, id int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("work %d finish\n", id)
					return
				default:
					fmt.Printf("work %d run\n", id)
					time.Sleep(1 * time.Second)
				}
			}
		}(ctx, i)
	}

	time.Sleep(1 * time.Microsecond)
	cancel()
	time.Sleep(1 * time.Second)
}

func withTimeout() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("work timeout")
				return
			default:
				fmt.Println("work run")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
	_ = cancel
	wg.Wait()
}

func withDeadline() {
	ctx := context.Background()
	t := time.Now()
	//设定超时时间为5s后
	t = t.Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, t)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("work timeout")
				return
			default:
				fmt.Println("work run")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
	_ = cancel
	wg.Wait()
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " work finish")
			return
		default:
			fmt.Println(name, " work run")
			time.Sleep(1 * time.Second)
		}
	}
}

func childCancelCtx() {
	ctx := context.Background()
	ctx1, cancel1 := context.WithCancel(ctx)
	ctx2, _ := context.WithCancel(ctx1)
	ctx3, _ := context.WithCancel(ctx2)

	/*
		关闭 cancel1 他会逐级往下关闭 ctx1,ctx2,ctx3 都会关闭
	*/
	go work(ctx1, "ctx1")
	go work(ctx2, "ctx2")
	go work(ctx3, "ctx3")

	time.Sleep(500 * time.Millisecond)
	cancel1()
	time.Sleep(1 * time.Second)
}

func childCancelCtx2() {
	ctx := context.Background()
	ctx1, _ := context.WithCancel(ctx)
	ctx2, cancel2 := context.WithCancel(ctx1)
	ctx3, _ := context.WithCancel(ctx2)

	/*
		关闭 cancel2 他会逐级往下关闭 ctx2,ctx3 会关闭,ctx1 不会被关闭
	*/
	go work(ctx1, "ctx1")
	go work(ctx2, "ctx2")
	go work(ctx3, "ctx3")

	time.Sleep(500 * time.Millisecond)
	cancel2()
	time.Sleep(1 * time.Second)
}
