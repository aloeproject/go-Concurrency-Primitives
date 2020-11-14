package once

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type Once struct {
	sync.Once
}

/*
判断是否经过初始化过
*/
func (o *Once) Done() bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(&o))) == 1
}
