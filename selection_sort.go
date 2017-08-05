package demo

func SelectionSort(a []int) []int {
	for i := 0; i < len(a); i++ { // n
		smallest := a[i]              // 1
		index := i                    // 1
		for j := i; j < len(a); j++ { //n
			if a[j] < smallest { //1
				smallest = a[j] //1
				index = j       //1
			}
		}
		tmp := a[i]     // 1
		a[i] = a[index] // 1
		a[index] = tmp  // 1
	}
	return a
}

// loop invariant:
//	for each i in [0,n], slice[0:i] is ordered ascending
// n ( 2 + n ( 1 + [0|2] ) + 3)
// -> theta(n^2)
