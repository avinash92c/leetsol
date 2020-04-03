
func minTimeToVisitAllPoints(points [][]int) int {
	count := 0
	for i := 0; i < len(points)-1; i++ {
		x1, y1, x2, y2 := points[i][0], points[i][1], points[i+1][0], points[i+1][1]
		//EQUATE X SIDES FIRST, THEN Y
        d1 := int(math.Abs(float64(x1 - x2)))
        d2 :=int( math.Abs(float64(y1 - y2)))
		if d1 == d2 {
			count += d1
		} else if d1 < d2 {
			count += d2
		} else {
			count += d1
		}
	}
	return count
}
