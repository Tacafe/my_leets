/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
    if root == nil {
        return 0
    }
    max := 0
    _ = helper(root, &max)
    return max
}

func helper(node *TreeNode, max *int) int {
    v, summary := 0, 0
    if node.Left != nil {
        vl := helper(node.Left, max)
        if node.Val == node.Left.Val {
            v = vl + 1
            summary += vl + 1
        }
    }
    if node.Right != nil {
        vr := helper(node.Right, max)
        if node.Val == node.Right.Val {
            if vr + 1 > v {
                v = vr + 1
            }
            summary += vr + 1
        }
    }
    if summary > *max {
        *max = summary
    }
    return v
}

// @lc code=end

func main() {
	l := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	r := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	root := &TreeNode{
		Val:   5,
		Left:  l,
		Right: r,
	}

	longestUnivaluePath(root)
}
