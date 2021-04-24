package rate1

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"time"
)

func rate1() {
	//令牌桶 一个5个 每秒放1个进行恢复
	r := rate.NewLimiter(1, 5)
	ctx := context.Background()
	for {
		//r.Wait()
		//每次消耗2个
		err := r.WaitN(ctx, 2)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(time.Now())

		time.Sleep(1 * time.Second)
	}
}

//如果超过令牌桶的数量限制返回false
func allow1() {
	r := rate.NewLimiter(1, 5)
	for {
		//allow
		//allowN
		if !r.AllowN(time.Now(), 2) {
			fmt.Println("to many request")
		} else {
			fmt.Println("ok")
		}
		time.Sleep(1 * time.Second)
	}
}

/*
把限流加入到中间件中
*/
func weblimit() {
	r := rate.NewLimiter(1, 5)
	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("middleware!\n"))
			next.ServeHTTP(writer, request)
		})
	}
	myLimit := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if !r.Allow() {
				http.Error(writer, "to many request", http.StatusForbidden)
				return
			}
			next.ServeHTTP(writer, request)
		})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})
	//增加中间键
	http.ListenAndServe(":8080", middleware(myLimit(mux)))
}
