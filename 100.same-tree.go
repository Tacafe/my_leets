/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func depthTraversal(n *TreeNode, l *[]int) {
	if n == nil {
		*l = append(*l, 0)
		return
	}

	depthTraversal(n.Left, l)
	depthTraversal(n.Right, l)
	*l = append(*l, n.Val)
}

func isEqual(p []int, q []int) bool {
	if len(p) != len(q) {
		return false
	}
	for i, p := range p {
		if p != q[i] {
			return false
		}
	}
	return true
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	p_values := []int{}
	q_values := []int{}

	depthTraversal(p, &p_values)
	depthTraversal(q, &q_values)

	return isEqual(p_values, q_values)
}

// @lc code=end

