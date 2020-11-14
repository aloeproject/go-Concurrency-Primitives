package cond

import (
	"fmt"
	"sync"
)

type LimitedQueue struct {
	queue    []int
	size     int
	length   int
	lock     *sync.Mutex
	pushCond *sync.Cond
	popCond  *sync.Cond
	closed   bool
}

func NewLimitedQueue(size int) *LimitedQueue {
	lock := new(sync.Mutex)
	cond1 := sync.NewCond(new(sync.Mutex))
	cond2 := sync.NewCond(new(sync.Mutex))
	return &LimitedQueue{
		queue:    make([]int, 0),
		size:     size,
		lock:     lock,
		length:   0,
		pushCond: cond1,
		popCond:  cond2,
		closed:   false,
	}
}

func (q *LimitedQueue) add(val int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, val)
	q.length++
}

/**
如果元素满了会被阻塞
*/
func (q *LimitedQueue) Push(val int) error {
	if q.size == 0 {
		return fmt.Errorf("push fail size is empty")
	}
	q.popCond.L.Lock()
	defer q.popCond.L.Unlock()

	if q.isClose() == true {
		return fmt.Errorf("queue is close")
	}

	for q.getLength() == q.size {
		//产生一个阻塞
		q.popCond.Wait()
		//发送一个广播
		q.pushCond.Broadcast()
	}

	if q.length < q.size {
		q.add(val)
	}

	return nil
}

func (q *LimitedQueue) pop() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	cur := q.queue[0]
	q.queue = q.queue[1:]
	q.length--
	return cur
}

func (q *LimitedQueue) getLength() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.length
}

/*
如果没有元素了则柱塞
如果被关闭了，且没有元素了释放阻塞
*/
func (q *LimitedQueue) Pop() (int, error) {
	q.pushCond.L.Lock()
	defer q.pushCond.L.Unlock()

	//如果为空等待被唤醒
	for q.getLength() == 0 {
		if q.isClose() == true {
			return 0, fmt.Errorf("queue is close")
		}
		q.pushCond.Wait()
	}

	cur := q.pop()
	//释放push的柱塞
	q.popCond.Broadcast()
	return cur, nil
}

func (q *LimitedQueue) isClose() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.closed
}

func (q *LimitedQueue) Close() {
	q.lock.Lock()
	q.closed = true
	//防止被阻塞
	q.pushCond.Broadcast()
	q.lock.Unlock()
}
