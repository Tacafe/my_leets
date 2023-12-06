/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {
	lTreeNodes := [][]int{{-200}}
	rTreeNodes := [][]int{{-200}}

	if root.Left != nil {
		lTreeNodes[0] = make([]int, root.Left.Val)
		getSubtree(lTreeNodes, root.Left, 1)
	}

	if root.Right != nil {
		rTreeNodes[0] = make([]int, root.Right.Val)
		getSubtree(rTreeNodes, root.Right, 1)
	}

	return check(lTreeNodes, rTreeNodes)
}

func getSubtree(nodesInEachLevels [][]int, n *TreeNode, level int) [][]int {
	if len(nodesInEachLevels) < level + 1 {
		nodesInEachLevels = append(*nodesInEachLevels, []int{-200, -200})
	}

	if n.Left != nil {
		nodesInEachLevels[level][0] = n.Left.Val
		getSubtree(&nodesInEachLevels, n.Left, level+1)
	}

	if n.Right != nil {
		nodesInEachLevels[level][1] = n.Right.Val
		getSubtree(&nodesInEachLevels, n.Right, level+1)
	}
}

func check(lTree *[][]int, rTree *[][]int) bool {
	// depth
	if len(lTree) != len(rTree) {
		return false
	}

	fmt.Println(lTree)
	fmt.Println(rTree)
	return false
}

// @lc code=end

/*
memo:
rootから見た、left nodesとright nodesを何らかの形で表現できれば良い

対称
[[2], [3,4]]
[[2], [4,3]]

これはだめ
[[2], [null, 3]]
[[2], [null, 3]]

*/
