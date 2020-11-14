package syncmap

import (
	"sync"
	"testing"
)

/*
从下面的性能对比上看
写性能 RWMap_Set > SyncMap_Set
读性能 SyncMap_Get > RWMap_Get
 */


func BenchmarkRWMap_Set(b *testing.B) {
	obj := NewRWMap(0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		obj.Set(i, i)
	}
	b.StopTimer()
}

func BenchmarkRWMap_Get(b *testing.B) {
	obj := NewRWMap(0)
	for i := 0; i < 1000; i++ {
		obj.Set(i, i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		obj.Get(i)
	}
	b.StopTimer()
}

func BenchmarkSyncMap_Set(b *testing.B) {
	m := new(sync.Map)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.StopTimer()
}

func BenchmarkSyncMap_Get(b *testing.B) {
	m := new(sync.Map)
	for i := 0; i < 1000; i++ {
		m.Store(i, i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Load(i)
	}
	b.StopTimer()
}
