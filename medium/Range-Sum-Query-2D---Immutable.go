type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{matrix: matrix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    var sum int
	i := row1
	j := col1
    for i <= row2 {
		j = col1
		for j <= col2 {
			sum += this.matrix[i][j]
			j++
		}
		i++
		if i > row2 {
			break
		}
	}
	return sum
}




/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
