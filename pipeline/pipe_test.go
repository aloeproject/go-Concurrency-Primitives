package pipeline

import (
	"fmt"
	"testing"
)

func Test_pipe(t *testing.T) {
	//fmt.Println(runtime.NumCPU())
	done := make(chan struct{})
	defer close(done)
	val := create(done, 1, 2, 3)
	for v := range multiply(done, add(done, val, 1), 2) {
		fmt.Println(v)
	}
}
