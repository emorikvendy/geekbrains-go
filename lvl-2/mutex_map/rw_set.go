package mutex_map

import "sync"

type RWSet struct {
	mx sync.RWMutex
	m  map[int]float64
}

func NewRWSet() *RWSet {
	return &RWSet{
		m: map[int]float64{},
	}
}
func (set *RWSet) Get(key int) (float64, bool) {
	set.mx.RLock()
	defer set.mx.RUnlock()
	val, ok := set.m[key]
	return val, ok
}

func (set *RWSet) Has(key int) bool {
	set.mx.RLock()
	defer set.mx.RUnlock()
	_, ok := set.m[key]
	return ok
}

func (set *RWSet) Add(key int, value float64) {
	set.mx.Lock()
	defer set.mx.Unlock()
	set.m[key] = value
}
