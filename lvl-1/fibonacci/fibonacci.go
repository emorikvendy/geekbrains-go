package fibonacci

var fibonacci = map[int64]int64{0: 0, 1: 1, 2: 1}

func Loop(n int64) int64 {
	if n == 1 || n == 2 {
		return 1
	}
	if n == 0 {
		return 0
	}
	sign := int64(1)
	if n < 0 {
		sign = -sign
		n = -n
	}
	first, second := int64(1), int64(1)
	for i := int64(2); i < n; i++ {
		second, first = first, first+second
	}
	return sign * first
}

func Recursive(n int64) int64 {
	if n == 1 || n == 2 {
		return 1
	} else if n == 0 {
		return 0
	}
	if n < 0 {
		return -Recursive(-n)
	}

	return Recursive(n-1) + Recursive(n-2)
}

func RecursiveWithMap(n int64) int64 {
	if n < 0 {
		return -RecursiveWithMap(-n)
	}
	if value, ok := fibonacci[n]; ok {
		return value
	}
	fibonacci[n] = RecursiveWithMap(n-2) + RecursiveWithMap(n-1)

	return fibonacci[n]
}
