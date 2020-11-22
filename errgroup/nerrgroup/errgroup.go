package nerrgroup

import (
	"context"
	"fmt"
	"github.com/neilotoole/errgroup"
	"time"
)

func maxGoroutine() {
	//最大任务数，最大并行数
	numG, qSize := 2, 4
	ctx := context.Background()
	g, ctx := errgroup.WithContextN(ctx, numG, qSize)
	g.Go(func() error {
		taskName := "task1"
		fmt.Println(taskName)
		time.Sleep(2 * time.Second)
		return nil
	})

	g.Go(func() error {
		taskName := "task2"
		fmt.Println(taskName)
		time.Sleep(2 * time.Second)
		return nil
	})

	g.Go(func() error {
		taskName := "task3"
		fmt.Println(taskName)
		time.Sleep(2 * time.Second)
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("finish")
}
