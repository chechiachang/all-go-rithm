package algorithm

func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				tmp := a[j]
				a[j] = a[j-1]
				a[j-1] = tmp
			}
		}
	}
	return a
}
