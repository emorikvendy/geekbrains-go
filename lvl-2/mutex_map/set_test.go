package mutex_map

import (
	"math/rand"
	"testing"
)

func BenchmarkSet9R1W(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				n := rand.Int() % 10
				key := rand.Int() % 100
				if n%10 == 0 {
					value := rand.Float64()
					set.Add(key, value)
				} else {
					_ = set.Has(key)
				}
			}
		})
	})
}

func BenchmarkRWSet9R1W(b *testing.B) {
	var set = NewRWSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				n := rand.Int() % 10
				key := rand.Int() % 100
				if n%10 == 0 {
					value := rand.Float64()
					set.Add(key, value)
				} else {
					_ = set.Has(key)
				}
			}
		})
	})
}

func BenchmarkSet9W1R(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				n := rand.Int() % 10
				key := rand.Int() % 100
				if n%10 == 0 {
					_ = set.Has(key)
				} else {
					value := rand.Float64()
					set.Add(key, value)
				}
			}
		})
	})
}

func BenchmarkRWSet9W1R(b *testing.B) {
	var set = NewRWSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				n := rand.Int() % 10
				key := rand.Int() % 100
				if n%10 == 0 {
					_ = set.Has(key)
				} else {
					value := rand.Float64()
					set.Add(key, value)
				}
			}
		})
	})
}

func BenchmarkSet5W5R(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				n := rand.Int() % 10
				key := rand.Int() % 100
				if n%2 == 0 {
					value := rand.Float64()
					set.Add(key, value)
				} else {
					_ = set.Has(key)
				}
			}
		})
	})
}

func BenchmarkRWSet5W5R(b *testing.B) {
	var set = NewRWSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				n := rand.Int() % 10
				key := rand.Int() % 100
				if n%2 == 0 {
					value := rand.Float64()
					set.Add(key, value)
				} else {
					_ = set.Has(key)
				}
			}
		})
	})
}
