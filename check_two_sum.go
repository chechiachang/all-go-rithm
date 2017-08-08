package algorithm

//TODO incorrect
func CheckTwoSum(a []int, b int) bool {
	for i := 0; i < len(a)-1; i++ { // theta(n)
		isThere := BinearySearch(a[i+1:len(a)], b-a[i])
		if isThere > 0 {
			return true
		}
	}

	return false
}
