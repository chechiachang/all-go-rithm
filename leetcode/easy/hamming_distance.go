package main

import (
	"strconv"
)

func HammingDistance(x int, y int) int {
	if x == y {
		return 0
	} else {
		xs := strconv.FormatInt(int64(x), 2)
		ys := strconv.FormatInt(int64(y), 2)
		count := 0
		if x > y {
			for i, _ := range ys {
				if xs[i] != ys[i] {
					count++
				}
			}
			count += len(xs) - len(ys)
		} else {
			for i, _ := range xs {
				if xs[i] != ys[i] {
					count++
				}
			}
			count += len(ys) - len(xs)
		}
		return count
	}
}
