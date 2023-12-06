/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < len(dp); i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}

// @lc code=end
