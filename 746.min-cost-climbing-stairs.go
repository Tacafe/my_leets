/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] = min(cost[i]+cost[i-1], cost[i]+cost[i-2])
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}

// @lc code=end

