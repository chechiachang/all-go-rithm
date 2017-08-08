package algorithm

func BinearySearch(a []int, b int) int {
	if a[len(a)/2] == b {
		return len(a) / 2

	} else if a[len(a)/2] < b {
		return BinearySearch(a[0:len(a)/2], b)

	} else {
		return len(a)/2 + BinearySearch(a[len(a)/2:len(a)], b)
	}
}

/**

c	1 2 3 4 5 6 7 8 9 10

c	1 2 3 4 5

c	1 2 3

c	1 2

c	1

height = log(n)
**/
