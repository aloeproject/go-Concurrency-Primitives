package pool

import (
	"sync"
)

func buf() {
	bp := sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}
	by := bp.Get().([]byte)
	by[0] = 'a'
}
