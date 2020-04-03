func countServers(grid [][]int) int {
	//COLLECT X AND Y INTO SEPERATE SLICES
	count := 0
	for i, val := range grid {
		for j, _ := range val {
			// SCAN SIDE WAYS IN SAME ARRAY
			found := false
			for k := 0; k < len(val); k++ {
				if k != j && val[k] == 1 && val[j] == 1 {
					found = true
					break
				}
			}
			if !found {
				for x := 0; x < len(grid); x++ {
					if grid[x][j] == 1 && grid[i][j] == 1 && x != i {
						found = true
						break
					}
				}
			}
			// fmt.Println(val[j])
			// fmt.Println(found)
			if found {
				count++
			}
		}
	}
	return count
}
