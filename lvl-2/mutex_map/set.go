package mutex_map

import "sync"

type Set struct {
	mx sync.Mutex
	m  map[int]float64
}

func NewSet() *Set {
	return &Set{
		m: map[int]float64{},
	}
}

func (set *Set) Get(key int) (float64, bool) {
	set.mx.Lock()
	defer set.mx.Unlock()
	val, ok := set.m[key]
	return val, ok
}

func (set *Set) Has(key int) bool {
	set.mx.Lock()
	defer set.mx.Unlock()
	_, ok := set.m[key]
	return ok
}

func (set *Set) Add(key int, value float64) {
	set.mx.Lock()
	defer set.mx.Unlock()
	set.m[key] = value
}
