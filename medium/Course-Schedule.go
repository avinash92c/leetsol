type Node struct {
	Val       int
	Connected []*Node
	Visited   bool
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	nodes := prepareGraph(prerequisites, numCourses)
	for _, val := range nodes {
		// fmt.Println("Fresh GO")
		if len(val.Connected) > 0 {
			if !scan(val, 0, numCourses) {
				return false
			}
		}
		// resetVisited(nodes)
	}
	return true
}

func resetVisited(nodes []*Node) {
	for _, val := range nodes {
		val.Visited = false
	}
}

func scan(node *Node, v, n int) bool {
	v++

	// fmt.Println("Nodes - " + strconv.Itoa(n) + " - V - " + strconv.Itoa(v) + " - Curr - " + strconv.Itoa(node.Val))

	if v > n {
		return false
	}
	for _, cnode := range node.Connected {
		if !scan(cnode, v, n) {
			return false
		}
	}
	return true
}

func prepareGraph(paths [][]int, n int) []*Node {
	nodes := make([]*Node, n)
	for idx := 0; idx < n; idx++ {
		nodes[idx] = &Node{Val: idx , Connected: make([]*Node, 0), Visited: false}
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
