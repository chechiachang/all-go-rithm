package algorithm

var (
	T1 = []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	A1 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	rand  = []int{9, 1, 8, 2, 7, 3, 6, 4, 5, 0}
	rand2 = []int{1, 5, 2, 6, 3, 7, 4, 8, 5, 9, 6, 0}
	a     = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	M1 = []int{1, 3, 5, 7, 9, 0, 2, 4, 6, 8}

	m1 = []int{1, 3, 5, 7, 9}
	m2 = []int{0, 2, 4, 6, 8}

	empty1 = []int{}
	empty2 = []int{}
	empty3 = []int{}

	bit1 = []int{1, 0, 0, 0, 1}
	bit2 = []int{0, 1, 1, 1, 0}
	bit3 = []int{1, 1, 1, 1, 1}
)

func Equals(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
