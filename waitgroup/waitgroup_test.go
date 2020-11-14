package waitgroup

import (
	"fmt"
	"testing"
)

func TestWaitGroup_WaiterCount(t *testing.T) {
	wg := new(WaitGroup)
	wg.Add(2)
	ct, ct2 := wg.WaiterCount()

	fmt.Println(ct, ct2)
}
