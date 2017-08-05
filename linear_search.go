package demo

func LinearSearch(array []int, v int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == v {
			return i
		}
	}
	return 0
}

/**
loop invariant:
	for i in [0, n], all element in slice[0:i] is not equals to v

Initialization:
	i = 0


Worse case:
	v = i[n]
	time = n

Best case:
	v = i[0]
	time = 1

Average case:
	v = i[n/2]
	time = n/2
**/
