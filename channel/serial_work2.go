package channel

import "sync"

type JopWorkV2 struct {
	w       []Work
	mutex   *sync.Mutex
	group   *sync.WaitGroup
	ch      chan struct{}
	workNum int
}

func NewJopWorkV2(workNum int) *JopWorkV2 {
	j := &JopWorkV2{
		w:       make([]Work, 0),
		mutex:   new(sync.Mutex),
		group:   new(sync.WaitGroup),
		ch:      make(chan struct{}),
		workNum: workNum,
	}
	go j.run()
	return j
}

func (j *JopWorkV2) Submit(f Work) {
	j.group.Add(1)
	j.Add(f)
	j.ch <- struct{}{}
}

func (j *JopWorkV2) Add(f Work) {
	j.mutex.Lock()
	j.w = append(j.w, f)
	j.mutex.Unlock()
}

func (j *JopWorkV2) Pop() Work {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	job := j.w[0]
	j.w = j.w[1:]
	return job
}

func (j *JopWorkV2) run() {
	for i := 0; i < j.workNum; i++ {
		go func() {
			for _ = range j.ch {
				j.Pop()()
				j.group.Done()
			}
		}()
	}
}

func (j *JopWorkV2) Wait() {
	j.group.Wait()
	close(j.ch)
}
