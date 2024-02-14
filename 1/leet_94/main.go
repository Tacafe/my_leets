package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func innerTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	innerTraversal(node.Left, values)
	*values = append(*values, node.Val)
	innerTraversal(node.Right, values)
}

func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	innerTraversal(root, &ret)
	return ret
}

func main() {
	fmt.Println("Hello")
}
