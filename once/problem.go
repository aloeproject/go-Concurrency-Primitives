package once

import (
	"fmt"
	"net"
	"sync"
)

func onceErr1() {
	/*
		死锁
		外层的锁没有释放，在内层再加一层
	*/
	var o1 sync.Once
	o1.Do(func() {
		o1.Do(func() {
			fmt.Print("init")
		})
	})
}

func onceErr2() {
	/*
		下面就存在如果，连接到网址失败
		但是初始化只会一次，下次再次调用不会再起作用
		造成 使用的是 conn的空连接
	*/
	var once sync.Once
	var conn net.Conn
	once.Do(func() {
		conn, _ = net.Dial("tcp", "baidu.com")
	})
	conn.Write([]byte("33"))
}
