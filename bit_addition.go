package algorithm

func BitAddition(a []int, b []int) []int {
	c := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] == 1 && b[i] == 1 {
			c[i] = 0
			c[i+1]++
		} else {
			c[i] += a[i] + b[i]
		}
	}
	return c
}
