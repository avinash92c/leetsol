func countNegatives(grid [][]int) int {
	//LOOP INNER ARRAYS IN REVERSE
	count := 0
	for _, val := range grid {
		// CHECK FIRST ELEMENT IN ARRAY
		//  IF YES ADD LENGTH OF ARRAY TO COUNT
		mid := len(val) / 2
		if val[0] < 0 {
			count += len(val)
		} else if len(val) > mid+1 && val[mid+1] < 0 && val[mid] >= 0 {
			//CHECK MID +1 ELEMENT
			// IF -ve add length of subarray
			count += len(val[mid+1:])
		} else {
			for i := len(val) - 1; i >= 0; i-- {
				if val[i] < 0 {
					count += 1
				}
			}
		}
	}
	return count
}
