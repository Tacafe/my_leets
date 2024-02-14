/*
 * @lc app=leetcode id=1103 lang=golang
 *
 * [1103] Distribute Candies to People
 */

// @lc code=start
func distributeCandies(candies int, num_people int) []int {
	dist := make([]int, num_people)
	giveCount := 0

	for {
		for i := 0; i < num_people; i++ {
			giveCount += 1
			if candies < giveCount {
				giveCount = candies
			}

			dist[i] += giveCount
			candies -= giveCount

			if candies == 0 {
				return dist
			}
		}
	}
	return dist
}
// @lc code=end

