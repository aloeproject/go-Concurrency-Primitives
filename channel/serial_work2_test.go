package channel

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestJopWorkV2_Submit(t *testing.T) {
	j := NewJopWorkV2(4)
	var num int32
	for i := 0; i < 10; i++ {
		j.Submit(func() {
			atomic.AddInt32(&num, 1)
			fmt.Println(atomic.LoadInt32(&num))
			time.Sleep(1 * time.Second)
		})
	}
	fmt.Println("done")
	j.Wait()
}