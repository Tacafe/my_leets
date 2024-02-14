/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
func cloneGraph(node *Node) *Node {
	nodes := make(map[int]*Node)
	return fillNodes(node, nodes)
}

func fillNodes(node *Node, nodes map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	cp := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	nodes[node.Val] = cp
	for _, n := range node.Neighbors {
		_, ok := nodes[n.Val]
		if !ok {
			fillNodes(n, nodes)
		}
		cp.Neighbors = append(cp.Neighbors, nodes[n.Val])
	}
	return cp
}
// @lc code=end

