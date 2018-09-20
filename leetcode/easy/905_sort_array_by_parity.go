package easy

import (
//"fmt"
)

func sortArrayByParity(A []int) []int {
	i := 0
	j := len(A) - 1
	tmp := 0

	for i < j {
		// if A[i] is odd and A[i] is even, switch
		if A[i]%2 > A[j]%2 {
			tmp = A[i]
			A[i] = A[j]
			A[j] = tmp
		}

		// move front index
		if A[i]%2 == 0 {
			i++
		}

		// move back index
		if A[j]%2 == 1 {
			j--
		}
	}

	return A
}
