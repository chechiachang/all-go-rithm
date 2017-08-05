package demo

func Merge(a []int, p, q, r int) {
	b := make([]int, len(a))

	li := 0
	ri := 0

	for li+ri < r {
		if a[li] < a[q+ri] {
			b[li+ri] = a[li]
			li++
		} else {
			b[li+ri] = a[q+ri]
			ri++
		}
		//fmt.Println(b)
	}
	a = b
}

func MergeSort(a []int, p, r int) []int {
	if p < r {
		q := (p + r) / 2
		MergeSort(a, p, q)
		MergeSort(a, q, r)
		Merge(a, p, q, r)
	}
	return a
}

func Merge2(a []int, b []int) []int {
	c := make([]int, len(a)+len(b))

	ai := 0
	bi := 0

	for i := 0; i < len(c); i++ {
		if ai == len(a) {
			c[i] = b[bi]

		} else if bi == len(b) {
			c[i] = a[ai]

		} else if a[ai] < b[bi] {
			c[i] = a[ai]
			ai++
		} else {
			c[i] = b[bi]
			bi++
		}
		//fmt.Println(c)
	}
	return c
}

func MergeSort2(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	left := a[0 : len(a)/2]
	right := a[len(a)/2 : len(a)]
	left = MergeSort2(left)
	right = MergeSort2(right)

	return Merge2(left, right)
}
