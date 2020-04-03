type Node struct {
	Val       int
	Connected []*Node
	Visited   bool
}

var used []int

func findOrder(numCourses int, prerequisites [][]int) []int {
	used = make([]int, numCourses)
	nodes := prepareGraph(prerequisites, numCourses)
	res := make([]int, 0)
	//SCAN FULL DEPTH
	for _, val := range nodes {
		if len(val.Connected) > 0 {
			// scanpath(val, &res)
			if !scanpath(val, &res, 0, numCourses) {
				return []int{}
			}
		} else {
			if used[val.Val] == 0 {
				res = append(res, val.Val)
			}
			used[val.Val] = 1
		}
	}
	return res
}

func scanpath(node *Node, res *[]int, v, n int) bool {
	v++
	if v > n {
		return false
	}
	if len(node.Connected) > 0 {
		for _, val := range node.Connected {
			// scanpath(val, res)
			if used[node.Val] == 0 && !scanpath(val, res, v, n) {
				return false
			}
		}
		if used[node.Val] == 0 {
			*res = append(*res, node.Val)
		}
		used[node.Val] = 1
	} else {
		if used[node.Val] == 0 {
			*res = append(*res, node.Val)
		}
		used[node.Val] = 1
	}
	return true
}

func prepareGraph(paths [][]int, n int) []*Node {
	nodes := make([]*Node, n)
	for idx := 0; idx < n; idx++ {
		nodes[idx] = &Node{Val: idx, Connected: make([]*Node, 0), Visited: false}
	}

	for _, val := range paths {
		a, b := val[0], val[1]
		// ga := nodes[a-1]
		// gb := nodes[b-1]
		ga := nodes[a]
		gb := nodes[b]
		ga.Connected = append(ga.Connected, gb)
		// gb.Connected = append(gb.Connected, ga)
	}
	return nodes
}
