package channel

import (
	"fmt"
	"time"
)

/*
下面代码会出现 goroutine 泄露
 */
func process(timeout time.Duration) bool {
	ch := make(chan bool)

	go func() {
		// 模拟处理耗时的业务
		time.Sleep(timeout + time.Second)
		/*
			在  <-time.After(timeout) 发送超时，函数退出
		    但是 没有 ch<-true  被消费而产生阻塞，goroutine会被夯住，而无法关闭

		   解决这个的办法是,变成一个缓存channel
			ch := make(chan bool,1)
		*/
		ch <- true // block

		fmt.Println("exit goroutine")
	}()
	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return false
	}
}
