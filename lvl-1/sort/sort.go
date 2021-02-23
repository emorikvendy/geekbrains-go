package sort

import (
	"sort"
)

// bubble sorts the int64 slice and returns a new slice with sorted elements
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

// insertion sort of the int64 slice and returns a new slice with sorted elements
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
					copy(sortableSlice[1:], sortableSlice[:i])
					sortableSlice[0] = tmp
				}
				break
			} else if sortableSlice[j] < tmp {
				if modify {
					copy(sortableSlice[j+1:i+1], sortableSlice[j:i])
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

// insertion sort of the int64 slice and returns a new slice with sorted elements
//
// Algorithm pseudocode:
//     for j = 2 to A.length do
//        key = A[j]
//        i = j-1
//        while (i > 0 and A[i] > key) do
//            A[i + 1] = A[i]
//            i = i - 1
//        end while
//        A[i+1] = key
//     end for[5]
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

// built-in sort of the int64 slice and returns a new slice with sorted elements
func BuiltIn(slice []int64) []int64 {
	sortableSlice := make([]int64, len(slice))
	copy(sortableSlice, slice)
	sort.Slice(sortableSlice, func(i, j int) bool { return sortableSlice[i] < sortableSlice[j] })
	return sortableSlice
}
