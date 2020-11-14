package waitgroup

import (
	"fmt"
	"sync"
	"testing"
)

func TestCounter_Count(t *testing.T) {
	wg := new(sync.WaitGroup)
	obj := new(Counter)
	wg.Add(1)
	worker(obj, wg)
	wg.Wait()

	fmt.Println(obj.Count())
}
