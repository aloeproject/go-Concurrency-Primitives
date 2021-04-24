package class4

import (
	"context"
	"fmt"
	"time"
)

func ctxExample() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case <-ctx.Done():
			//这里会返回 context canceled
			fmt.Println(ctx.Err())
		}
	}()
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
