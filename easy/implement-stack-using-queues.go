
type MyStack struct {
	Ds    *Node
	Dslen int
}

type Node struct {
	Val  int
	Prev *Node
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		Ds:    nil,
		Dslen: 0,
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
    node:= &Node{Val: x, Prev: nil}
	if this.Ds != nil {
		node.Prev = this.Ds
	} 
	this.Ds = node
	this.Dslen++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Ds != nil {
		retval := this.Ds.Val
		this.Ds = this.Ds.Prev
        this.Dslen--
		return retval
		
	}
	return 0
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Ds.Val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    
	return this.Dslen == 0
}
