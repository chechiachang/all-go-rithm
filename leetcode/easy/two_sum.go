package easy

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func twoSumHashmap(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		complement := target - v
		if _, ok := m[complement]; ok && m[complement] != i {
			return []int{i, m[complement]}
		}
	}

	return []int{0, 0}
}
