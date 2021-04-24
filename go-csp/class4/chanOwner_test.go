package class4

import (
	"fmt"
	"testing"
)

func TestChOwner(t *testing.T) {
	/*

	 */

	chanOwner := func() <-chan int {
		//在函数的词法范围内去实例化channel,这将结果写入处理范围，防止其他的goroutine写入它
		result := make(chan int)
		go func() {
			//生产者来关闭 channel
			defer close(result)
			for i := 0; i < 5; i++ {
				result <- i
			}
		}()
		return result
	}
	//约束传入的 channel result<-chan 只能读
	consumer := func(result <-chan int) {
		for item := range result {
			fmt.Println(item)
		}
		fmt.Println("done")
	}
	//消费者 只能读取
	ch := chanOwner()
	/*
		这么写也可以，随意启动多个消费者去进行消费
	*/
	consumer(ch)
}
