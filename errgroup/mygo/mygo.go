package mygo

import (
	"context"
	"fmt"
)

/*
实现 go run
让他可以 安全的退出，并且有超时控制
*/

type MyGo struct {
	data chan string
	stop chan struct{}
}

func NewMyGo() *MyGo {
	return &MyGo{
		data: make(chan string, 2),
		stop: make(chan struct{}),
	}
}

func (m *MyGo) Event(ctx context.Context, data string) error {
	select {
	case m.data <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (m *MyGo) Run() {
	for v := range m.data {
		fmt.Println(v)
	//	time.Sleep(1 * time.Second)
	}
	m.stop <- struct{}{}
}

func (m *MyGo) Shutdown(ctx context.Context) {
	close(m.data)
	select {
	case <-m.stop:
	case <-ctx.Done():
	}
}
