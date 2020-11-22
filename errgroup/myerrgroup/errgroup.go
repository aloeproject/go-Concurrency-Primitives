package myerrgroup

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

/*
自定义 errgroup
官方扩展库 ErrGroup 没有实现可以取消子任务的功能，
请你课下可以自己去实现一个子任务可取消的 ErrGroup。
并且goroutine发生panic，程序不会崩溃
*/

type ErrGroup struct {
	group  sync.WaitGroup
	cancel func()
	fun    func(ctx context.Context) error
	ctx    context.Context
	once   sync.Once
	err    error
}

func WithContext(ctx context.Context) *ErrGroup {
	ctx, cancel := context.WithCancel(ctx)
	return &ErrGroup{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (g *ErrGroup) Go(f func(ctx context.Context) error) {
	g.group.Add(1)
	go g.work(f)
}

func (g *ErrGroup) work(f func(ctx context.Context) error) {
	var err error
	ctx := g.ctx
	//如果初始化没有初始化ctx，则初始化
	if ctx == nil {
		ctx = context.Background()
	}
	defer func() {
		defer g.group.Done()

		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			//打印错误栈
			err = fmt.Errorf("errgroup: panic recovered: %s\n%s", r, buf)
		}

		if err != nil {
			//只打印第一次错误
			g.once.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
	err = f(ctx)
}

func (g *ErrGroup) Wait() error {
	g.group.Wait()
	/*
	这一次调用的目的？
	 */
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
