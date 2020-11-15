package channel

import (
	"fmt"
	"time"
)

/*
有一道经典的使用 Channel 进行任务编排的题，你可以尝试做一下：有四个 goroutine，
编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，
要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
*/

var ch1 chan struct{}
var ch2 chan struct{}
var ch3 chan struct{}
var ch4 chan struct{}

func work1() {
	<-ch1
	fmt.Print(1)
	ch2 <- struct{}{}
}

func work2() {
	<-ch2
	fmt.Print(2)
	ch3 <- struct{}{}
}

func work3() {
	<-ch3
	fmt.Print(3)
	ch4 <- struct{}{}
}

func work4() {
	<-ch4
	fmt.Print(4)
}

func handle() {
	ch1 = make(chan struct{})
	ch2 = make(chan struct{})
	ch3 = make(chan struct{})
	ch4 = make(chan struct{})

	for {
		select {
		case <-time.After(1 * time.Second):
			go work1()
			go work2()
			go work3()
			go work4()
			//fmt.Println(3)
			ch1 <- struct{}{}
			fmt.Println()
			//fmt.Println(2)
		}
	}
}
