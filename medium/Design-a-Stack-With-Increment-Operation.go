type CustomStack struct {
	Ds     []int
	Dslen  int
	DsSize int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		Ds:     make([]int, 0),
		Dslen:  0,
		DsSize: maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if this.DsSize > this.Dslen { //SIZE NOT REACHED
		this.Ds = append(this.Ds, x)
		this.Dslen++
	}
}

func (this *CustomStack) Pop() int {
	if this.Dslen > 0 {
		val := this.Ds[this.Dslen-1]
		this.Ds = this.Ds[:this.Dslen-1]
		this.Dslen--
		return val
	}
	return -1
}

func (this *CustomStack) Increment(k int, val int) {
	if k >= this.Dslen {
		//INCREMENT ALL
		this.incrementData(this.Dslen, val)
	} else {
		this.incrementData(k, val)
	}
}

func (this *CustomStack) incrementData(endidx, val int) {
    // fmt.Println(endidx)
    // fmt.Println(this.Dslen)
	for i := 0; i < endidx; i++ {
		this.Ds[i] = this.Ds[i] + val
	}
}
