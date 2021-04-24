package mygo

import (
	"context"
	"testing"
	"time"
)

func Test_MyGo_Event(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	gg := NewMyGo()
	go gg.Run()
	gg.Event(ctx, "test1")
	gg.Event(ctx, "test2")
	gg.Event(ctx, "test3")
	//gg.Shutdown(ctx)
}
