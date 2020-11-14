package waitgroup

import (
	"sync"
	"unsafe"
)

/*
type WaitGroup struct {
	// 避免复制使用的一个技巧，可以告诉vet工具违反了复制使用的规则
	noCopy noCopy
	// 64bit(8bytes)的值分成两段，高32bit是计数值，低32bit是waiter的计数
	// 另外32bit是用作信号量的
   // 因为64bit值的原子操作需要64bit对齐，但是32bit编译器不支持，所以数组中的元素在不同的架构中不一样，具体处理看下面的方法
   // 总之，会找到对齐的那64bit作为state，其余的32bit做信号量
	state1 [3]uint32

*/
type WaitGroup struct {
	sync.WaitGroup
}

func (wg *WaitGroup) WaiterCount() (uint32,uint32) {
	/*
	注意点：
	1. 这里用了一个函数来实现，更常见的可以自己封一个类。用函数实现时注意用指针传递wg
	2. 返回的两个值分别是state和wait，state是要完成的waiter计数值（即等待多少个goroutine完成）；wait是指有多少个sync.Wait在等待（和前面的waiter不是一个概念）。

	最后谈一点：
	nocopy这个字段之前一直没有注意，没想到还有这么巧妙的方法~
	 */
	var statep *uint64
	if uintptr(unsafe.Pointer(wg))%8 == 0 {
		statep = (*uint64)(unsafe.Pointer(wg))
	} else {
		statep = (*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(wg)) + unsafe.Sizeof(uint32(0))))
	}
	return uint32(*statep >> 32), uint32(*statep)
}
