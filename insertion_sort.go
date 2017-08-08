package algorithm

import "fmt"

func InsertionSort(A []int) []int {
	for i := 1; i < len(A); i++ { // n times
		for j := i; j > 0; j-- { // n ; n ; n
			if A[j] < A[j-1] { // 1 each for loop
				// 0 if false, 3 if true
				tmp := A[j]   // 1 each for loop
				A[j] = A[j-1] // 1 each for loop
				A[j-1] = tmp  // 1 each for loop
			}
		}
	}
	return A
}

// n*3n*(1 + [0|3])
// 3n^2 to 12n^2 -> theta(n^3)

func InsertionSortRecursive(a []int) {
	Insert(a[0], a[1:len(a)])
}
func Insert(b int, a []int) {
	if len(a) < 1 {
		a = append(a, b)
	} else if len(a) == 1 {
		a = append(a, b)
		if a[0] > b {
			tmp := a[1]
			a[1] = a[0]
			a[0] = tmp
		}
	} else {
		fmt.Println(a)
		Insert(a[1], a[1:len(a)])
	}
}

// 3,
// 3, 1
