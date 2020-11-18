package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestGWork_Handle(t *testing.T) {
	obj := NewGWork()
	obj.GoHandle(func() {
		fmt.Println("run 1")
		time.Sleep(6 * time.Second)
	})
	obj.GoHandle(func() {
		fmt.Println("run 2")
		time.Sleep(2 * time.Second)
	})
	obj.GoHandle(func() {
		fmt.Println("run 3")
		time.Sleep(2 * time.Second)
	})
	obj.Wait()
}
