/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	thresh := len(nums) / 2
	itemMap := make(map[int]int)

	for _, item := range nums {
		itemMap[item]++
		if itemMap[item] > thresh {
			return item
		}
	}
	return 0

}

// @lc code=end

