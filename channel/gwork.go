package channel

import (
	"log"
	"sync"
	"time"
)

type GWork struct {
	wg      *sync.WaitGroup
	timeout time.Duration
}

func NewGWork() *GWork {
	return &GWork{
		wg:      new(sync.WaitGroup),
		timeout: 5 * time.Second,
	}
}

func (g GWork) work(f func(), wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.NewTicker(g.timeout)
	//设定容量为1 防止goroutine泄露
	done := make(chan struct{}, 1)
	go func() {
		f()
		done <- struct{}{}
	}()
	select {
	case <-done:
		return
	case <-t.C:
		log.Print("time out")
		return
	}
}

func (g GWork) GoHandle(f func()) {
	g.wg.Add(1)
	go g.work(f, g.wg)
}

func (g GWork) Wait() {
	g.wg.Wait()
}
