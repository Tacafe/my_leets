/*
 * @lc app=leetcode id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
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

func getLeaves(node *TreeNode) []int {
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}

	leaves := []int{}
	if node.Left != nil {
		leaves = append(leaves, getLeaves(node.Left)...)
	}
	if node.Right != nil {
		leaves = append(leaves, getLeaves(node.Right)...)
	}
	return leaves
}

func isEqual(l1 []int, l2 []int) bool {
	if len(l1) != len(l2) { return false }
	for i := 0; i < len(l1); i ++ {
		if l1[i] != l2[i] { return false }
	}
	return true
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return isEqual(getLeaves(root1), getLeaves(root2))
}
// @lc code=end

