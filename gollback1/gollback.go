package gollback1

import (
	"context"
	"fmt"
	"github.com/vardius/gollback"
	"reflect"
	"time"
)

/*
它会等待所有的异步函数（AsyncFunc）都执行完才返回，而且返回结果的顺序和传入的函数的顺序保持一致。
第一个返回参数是子任务的执行结果，第二个参数是子任务执行时的错误信息
*/
func all() {
	ctx := context.Background()
	rs, errs := gollback.All(ctx, func(ctx context.Context) (i interface{}, e error) {
		return "2", nil
	}, func(ctx context.Context) (i interface{}, e error) {
		return 2, fmt.Errorf("error 2")
	}, func(ctx context.Context) (i interface{}, e error) {
		return 3, fmt.Errorf("error 3")
	})

	fmt.Println(rs)
	fmt.Println(reflect.TypeOf(rs))
	fmt.Println(errs)
}

/*
如果异步函数中一个没有错误就立刻返回，如果都有错误，返回最后一个错误
*/
func race() {
	ctx := context.Background()
	rs, errs := gollback.Race(ctx, func(ctx context.Context) (i interface{}, e error) {
		return 1, nil
	}, func(ctx context.Context) (i interface{}, e error) {
		return 2, nil
	}, func(ctx context.Context) (i interface{}, e error) {
		return 3, nil
	})

	fmt.Println(rs)
	fmt.Println(reflect.TypeOf(rs))
	fmt.Println(errs)
}

func race1() {
	ctx := context.Background()
	rs, errs := gollback.Race(ctx, func(ctx context.Context) (i interface{}, e error) {
		//return 1, nil
		return 1, fmt.Errorf("error 1")
	}, func(ctx context.Context) (i interface{}, e error) {
		return 2, fmt.Errorf("error 2")
	}, func(ctx context.Context) (i interface{}, e error) {
		time.Sleep(1 * time.Second)
		return 3, nil
	})

	fmt.Println(rs)
	fmt.Println(reflect.TypeOf(rs))
	fmt.Println(errs)
}

func retry() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	//如果发生错误进行n次重试，或者如果超时了也进行取消
	rs, err := gollback.Retry(ctx, 3, func(ctx context.Context) (i interface{}, e error) {
		fmt.Println("create error")
		return nil, fmt.Errorf("error")
	})
	fmt.Println(rs)
	fmt.Println(err)
}
