package medium

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rowLen := len(grid)
	colLen := len(grid[0])

	// Count miminum sum for first col
	//	1 2 3			1 3 6
	//	1 1 1			1 1 1
	//	1 2 3			1 2 3
	for i := 1; i < colLen; i++ {
		grid[0][i] += grid[0][i-1]
	}

	// Count miminum sum for first row
	//	1 2 3			1 3 6
	//	1 1 1			2 1 1
	//	1 2 3			3 2 3
	for j := 1; j < rowLen; j++ {
		grid[j][0] += grid[j-1][0]
	}

	// Count and choose minimun sum for inside path
	//	1 3 6
	//	2 1 1
	//	3 2 3
	for y := 1; y < rowLen; y++ {

		// count horizontally by row from left to right
		for x := 1; x < colLen; x++ {

			// Count min sum from left node
			left := grid[y][x-1]

			// Count min sum from top node
			up := grid[y-1][x]

			if left > up {
				grid[y][x] += up
			} else {
				grid[y][x] += left
			}
		}
	}

	//	1 3 6
	//	2 3 4
	//	3 5 7

	return grid[rowLen-1][colLen-1]
}
