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

/*
index 2からはじめて
[0, 1, 2*, 3, 4, 5]

戻る方向のcostのうち、少なくなる方を採用する
[0, 1, 2+0, 3, 4, 5]

indexをiterateさせると、それまでの累積コストがはんせいされていく
[0, 1, 2+0, 3+1, 4+(2+0), 5+(3+1)]
[0, 1, 2, 4, 6, 8]

最大2stepなので、末尾[6, 8]のうち、小さい方が最小コスト
*/


