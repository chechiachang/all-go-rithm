package hard

import (
//"fmt"
)

func LongestConsecutiveSequence(nums []int) int {

	h := &hashSet{
		set: make(map[int]bool),
	}

	if len(nums) == 0 {
		return 0
	}

	for _, n := range nums {
		h.add(n)
	}

	//fmt.Println(h)

	longest_streak := 0
	for _, n := range nums {
		// Has next
		if ok := h.get(n + 1); ok {

			current := n
			current_streak := 0

			for h.get(current + 1) {
				current++
				current_streak++
			}

			if current_streak > longest_streak {
				longest_streak = current_streak
			}
		}

	}

	return longest_streak + 1 //Add the first count
}

type hashSet struct {
	set map[int]bool
}

func (h *hashSet) add(i int) bool {
	_, found := h.set[i]
	h.set[i] = true
	return !found
}

func (h *hashSet) get(i int) bool {
	_, found := h.set[i]
	return found
}
