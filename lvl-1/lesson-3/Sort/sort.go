package sort

import (
	"sort"
)

func Bubble(slice []int64) []int64 {
	sortableSlice := make([]int64, len(slice))
	copy(sortableSlice, slice)
	for j := len(sortableSlice) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if sortableSlice[i] > sortableSlice[i+1] {
				sortableSlice[i+1], sortableSlice[i] = sortableSlice[i], sortableSlice[i+1]
			}
		}
	}
	return sortableSlice
}

func Insertion(slice []int64) []int64 {
	var tmp int64
	var j int
	sortableSlice := make([]int64, len(slice))
	copy(sortableSlice, slice)
	for i := 1; i < len(sortableSlice); i++ {
		tmp = sortableSlice[i]
		j = i
		modify := false
		for {
			j--
			if j < 0 {
				if modify {
					copy(sortableSlice, append([]int64{tmp}, sortableSlice[:i]...))
				}
				break
			} else if sortableSlice[j] < tmp {
				if modify {
					copy(sortableSlice, append(sortableSlice[:j+1], sortableSlice[j:i]...))
					sortableSlice[j+1] = tmp
				}
				break
			} else {
				modify = true
			}
		}
	}
	return sortableSlice
}

func InsertionWiki(slice []int64) []int64 {
	var tmp int64
	var j int
	sortableSlice := make([]int64, len(slice))
	copy(sortableSlice, slice)
	for i := 1; i < len(sortableSlice); i++ {
		tmp = sortableSlice[i]
		for j = i - 1; j >= 0 && sortableSlice[j] > tmp; j-- {
			sortableSlice[j+1] = sortableSlice[j]
		}
		sortableSlice[j+1] = tmp
	}
	return sortableSlice
}

func BuiltIn(slice []int64) []int64 {
	sortableSlice := make([]int64, len(slice))
	copy(sortableSlice, slice)
	sort.Slice(sortableSlice, func(i, j int) bool { return sortableSlice[i] < sortableSlice[j] })
	return sortableSlice
}
