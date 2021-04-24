package channel

import "sync"

type Work func()

type JopWork struct {
	w     []Work
	mutex *sync.Mutex
	group *sync.WaitGroup
	once  sync.Once
	ch    chan struct{}
}

func NewJopWork() *JopWork {
	j := &JopWork{
		w:     make([]Work, 0),
		mutex: new(sync.Mutex),
		group: new(sync.WaitGroup),
		ch:    make(chan struct{}),
	}
	go j.run()
	return j
}

func (j *JopWork) Submit(f Work) {
	j.mutex.Lock()
	j.w = append(j.w, f)
	j.group.Add(1)
	j.once.Do(func() {
		j.ch <- struct{}{}
	})
	j.mutex.Unlock()
}

func (j *JopWork) run() {
	for {
		select {
		case <-j.ch:
			go func() {
				defer func() {
					j.ch <- struct{}{}
					j.group.Done()
				}()
				job := j.w[0]
				j.w = j.w[1:]
				job()
			}()
		}
	}
}

func (j *JopWork) Wait() {
	j.group.Wait()
	close(j.ch)
}

