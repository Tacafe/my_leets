/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(helper(nums, 0, len(nums)-2), helper(nums, 1, len(nums)-1))
}

func helper(nums []int, l, r int) int {
	A, B := 0, 0
	for i := l; i <= r; i++ {
		a, b := A, B
		A = b + nums[i]
		B = max(a, b)
	}
	return max(A, B)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

