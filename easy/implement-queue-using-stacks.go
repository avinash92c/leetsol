
type MyQueue struct {
	Ds []int
	// Dslen int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		Ds: make([]int, 0),
		// Dslen: 0,
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Ds = append(this.Ds, x)
	// this.Dslen++
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	e := this.Ds[0]
	this.Ds = this.Ds[1:len(this.Ds)]
	// this.Dslen--
	return e
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.Ds[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.Ds) == 0 {
		return true
	}
	return false
}
