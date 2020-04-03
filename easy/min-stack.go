
type MinStack struct {
	Ds    []int
	Dslen int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Ds:    make([]int, 0),
		Dslen: 0,
	}
}

func (this *MinStack) Push(x int) {
	this.Ds = append(this.Ds, x)
	this.Dslen++
}

func (this *MinStack) Pop() {
	this.Ds = this.Ds[:this.Dslen-1]
	this.Dslen--
    fmt.Println(this.Ds)
}

func (this *MinStack) Top() int {
	return this.Ds[this.Dslen-1]
}

func (this *MinStack) GetMin() int {
    minVal := this.Ds[0]
	
	for _, val := range this.Ds {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}
