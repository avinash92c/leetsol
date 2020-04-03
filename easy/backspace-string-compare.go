func backspaceCompare(S string, T string) bool {
	return scandelete(S) == scandelete(T)
}

func scandelete(V string) string {
	m := Constructor()
	for i := 0; i < len(V); i++ {
		if string(V[i]) == "#" {
			m.Pop()
		} else {
			m.Push(string(V[i]))
		}
	}
	return m.Result()
}

type MinStack struct {
	Ds    []string
	Dslen int
}

func Constructor() MinStack {
	return MinStack{
		Ds:    make([]string, 0),
		Dslen: 0,
	}
}

func (this *MinStack) Result() string {
	val := ""
	for i := 0; i < len(this.Ds); i++ {
		val += this.Ds[i]
	}
	return val
}

func (this *MinStack) Peek() string {
	return this.Ds[this.Dslen-1]
}

func (this *MinStack) Push(x string) {
	this.Ds = append(this.Ds, x)
	this.Dslen++
}

func (this *MinStack) Pop() {
	if this.Dslen > 0 {
		this.Ds = this.Ds[:this.Dslen-1]
		this.Dslen--
	}
}
