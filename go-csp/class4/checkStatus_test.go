package class4

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCheckStatus(t *testing.T) {
	//错误应该不是在里面被打印，而是传递出来
	type Result struct {
		Error    error
		Response *http.Response
	}

	checkStatus := func(done chan struct{}, urls ...string) chan Result {
		res := make(chan Result)
		go func() {
			defer close(res)
			for _, url := range urls {
				var r Result
				resp, err := http.Get(url)
				r.Response = resp
				//协调多个 goroutine 时，应该制定更复杂的规则继续，来出来相应错误
				r.Error = err

				select {
				case res <- r:
				case <-done:
					return
				}
			}
		}()
		return res
	}

	done := make(chan struct{})
	defer close(done)
	for r := range checkStatus(done, "https://www.baidu.com", "http://aa1222222.com") {
		if r.Error != nil {
			fmt.Printf("error:%v\n", r.Error)
			continue
		}
		fmt.Printf("response %v\n", r.Response.Status)
	}

}
