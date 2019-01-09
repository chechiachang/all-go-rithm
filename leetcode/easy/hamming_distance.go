package easy

import (
	//"fmt"
	"strconv"
)

func HammingDistance(x int, y int) int {
	if x == y {
		return 0
	} else {
		count := 0
		for x != 0 || y != 0 {
			if x%2 != y%2 {
				count++
			}
			x = x >> 1
			y = y >> 1
			//fmt.Printf("x %v y %v count %v \n", x, y, count)
		}
		return count
	}
}

func HammingDistance2(x int, y int) int {
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
				//fmt.Printf("xs[i] %v ys[i] %v count %v \n", xs[i], ys[i], count)
			}
			count += len(xs) - len(ys)
			//fmt.Printf("(%v, %v) count %v xs %v ys %v \n", x, y, count, len(xs), len(ys))
		} else {
			for i, _ := range xs {
				if xs[i] != ys[i] {
					count++
				}
				//fmt.Printf("xs[i] %v ys[i] %v count %v \n", xs[i], ys[i], count)
			}
			count += len(ys) - len(xs)
			//fmt.Printf("(%v, %v) count %v xs %v ys %v \n", x, y, count, len(xs), len(ys))
		}
		return count
	}
}
