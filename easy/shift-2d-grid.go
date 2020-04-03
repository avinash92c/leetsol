// import "fmt"
func shiftGrid(grid [][]int, k int) [][]int {
	grid1d := make([]int, 0)
	x, y := len(grid), 0
	// BREAKDOWN GRID TO 1D ARRAY
	for _, val := range grid {
		y = len(val)
		for _, sval := range val {
			grid1d = append(grid1d, sval)
		}
	}

	// COUNT SHIFT INDEX
	sidx := 0
	if len(grid1d) < k {
		sidx = k % len(grid1d)
        sidx = len(grid1d) - sidx
	} else {
		sidx = len(grid1d) - k
	}
    // fmt.Println(sidx)

	if sidx > 0 {
		gridend := grid1d[sidx:]
		grid1d = grid1d[:sidx]
		gridend = append(gridend, grid1d...)

		// GENERATE FINAL ARRAY
		res := make([][]int, 0)
		bidx := 0
		for i := 0; i < x; i++ {
			res = append(res, gridend[bidx:bidx+y])
			bidx += y
		}
		return res
	} else if sidx == 0 {
		return grid
	}
	return grid
}
