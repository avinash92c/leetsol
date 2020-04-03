func simplifyPath(path string) string {
	m := Constructor()

	//SPLIT STRING BY /
	pathparts := strings.Split(path, "/")

	for _, part := range pathparts {
		fmt.Println(part)
		if part == ".." {
			m.Pop()
		} else if part == "." || part == "./" {
			//IGNORE
		} else {
			if len(strings.Trim(part, "")) > 0 {
				// val := strings.ReplaceAll(part, "/", "")
				m.Push(part)
			}
		}
	}

	return m.GetPath()
}

type MyStack struct {
	Ds    *Node
	Dslen int
}

type Node struct {
	Val  string
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
func (this *MyStack) Push(x string) {
	node := &Node{Val: x, Prev: this.Ds}
	this.Ds = node
	this.Dslen++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() string {
	if this.Ds != nil {
		retval := this.Ds.Val
		this.Ds = this.Ds.Prev
		this.Dslen--
		return retval
	}
	return ""
}

func (this *MyStack) GetPath() string {
	res := ""
	currnode := this.Ds
	for currnode != nil {
		res = "/" + currnode.Val + res
		currnode = currnode.Prev
	}
    if len(res)==0{
        return  "/"
    }
	return res
}

