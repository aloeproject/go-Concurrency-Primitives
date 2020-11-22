package berrgroup

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/sync/errgroup"
	"sync"
	"time"
)

/*
因为原生的errgroup 不支持最大的设定goroutine上限，所以增加了上限的功能


有一点不太好的地方就是，一旦你设置了并发数，
超过并发数的子任务需要等到调用者调用 Wait 之后才会执行，
而不是只要 goroutine 空闲下来，就去执行。如果不注意这一点的话，
可能会出现子任务不能及时处理的情况，这是这个库可以优化的一点
*/

func maxGoroutine() {
	var g errgroup.Group
	g.GOMAXPROCS(2)
	var sum int
	var lock sync.Mutex
	for i := 0; i < 100; i++ {
		g.Go(func(ctx context.Context) error {
			lock.Lock()
			sum++
			lock.Unlock()
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
}

func errGroup1() {
	var g errgroup.Group
	g.GOMAXPROCS(2)
	//同样返回第一次出现的错误
	g.Go(func(ctx context.Context) error {
		taskName := "task1"
		fmt.Println(taskName)
		return nil
	})
	g.Go(func(ctx context.Context) error {
		taskName := "task2"
		fmt.Println(taskName)
		return fmt.Errorf("%s error", taskName)
	})
	g.Go(func(ctx context.Context) error {
		taskName := "task3"
		fmt.Println(taskName)
		return fmt.Errorf("%s error", taskName)
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

func panicDemo() {
	var g errgroup.Group
	g.GOMAXPROCS(2)
	//同样返回第一次出现的错误
	g.Go(func(ctx context.Context) error {
		taskName := "task1"
		fmt.Println(taskName)
		return nil
	})
	g.Go(func(ctx context.Context) error {
		taskName := "task2"
		fmt.Println(taskName)
		return fmt.Errorf("%s error", taskName)
	})
	/*
		panic捕获且，不会造成程序崩溃
		但是因为程序只会打印最早一次错误所以，后面的panic错误没有被发现
	*/
	g.Go(func(ctx context.Context) error {
		taskName := "task3"
		fmt.Println(taskName)
		panic(taskName + "panic")
		return fmt.Errorf("%s error", taskName)
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

func cancelWork() {
	ctx := context.Background()
	g := errgroup.WithCancel(ctx)
	g.GOMAXPROCS(3)
	//同样返回第一次出现的错误
	g.Go(func(ctx context.Context) error {
		name := "task1"
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, " cancel")
				return nil
			default:
				fmt.Println(name, "do some things")
				time.Sleep(3 * time.Second)
				//return nil
			}
		}
	})
	g.Go(func(ctx context.Context) error {
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

	/*
	task3 不会打印 do some things 最大goroutine只能2所以他没有被执行就会被取消了
	 */
	g.Go(func(ctx context.Context) error {
		name := "task3"
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, " cancel")
				return nil
			default:
				fmt.Println(name, "do some things")
				time.Sleep(3 * time.Second)
				//return nil
			}
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("finish")
	/*
	最终打印
	task1 do some things
	task2 do some things
	task3  cancel
	task1  cancel
	task2 error
	finish
	 */
}
