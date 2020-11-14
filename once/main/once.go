package main

import (
	"fmt"
	"sync"
)

func main() {
	o1 := new(sync.Once)

	o1.Do(func() {
		fmt.Println("first 1")
	})

	o1.Do(func() {
		fmt.Println("first 2")
	})
}
