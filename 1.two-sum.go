/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	ret := []int{}
	pairTargetMap := make(map[int]int)

	for i, num := range nums {
		pairTarget := target - num
		targetIdx, ok := pairTargetMap[num]
		if ok {
			return []int{i, targetIdx}
		} else {
			pairTargetMap[pairTarget] = i
		}
	}
	return ret
}
// @lc code=end

