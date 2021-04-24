package fanIn

import (
	"fmt"
	"runtime"
	"testing"
)

func Test_work(t *testing.T) {
	done := make(chan struct{})
	for v := range work(done, 1, 2, 3) {
		fmt.Println(v)
	}
}

func Test_fanIn(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	num := runtime.NumCPU()
	chArr := make([]<-chan int, 0)
	for i := 0; i < num; i++ {
		chArr = append(chArr, work(done, 1))
	}
	for i := range fanIn(done, chArr...) {
		fmt.Println(i)
	}
}


func Test_fanIn2(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	num := runtime.NumCPU()
	chArr := make([]<-chan int, 0)
	for i := 0; i < num; i++ {
		chArr = append(chArr, work2(done, 1))
	}
	for i := range fanIn(done, chArr...) {
		fmt.Println(i)
	}
}

func Test_fanIn4(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	num := runtime.NumCPU()
	chArr := make([]<-chan int, 0)
	for i := 0; i < num; i++ {
		chArr = append(chArr, nil)
	}
	for i := range fanOr(done, chArr...) {
		fmt.Println(i)
	}
}

func Test_fanIn3(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	num := runtime.NumCPU()
	chArr := make([]<-chan int, 0)
	for i := 0; i < num; i++ {
		chArr = append(chArr, work2(done, 1))
	}
	for i := range fanOr(done, chArr...) {
		fmt.Println(i)
	}
}