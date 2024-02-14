/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func hasItem(nums []int, target int) bool {
	for _, item := range nums {
		if item == target {
			return true
		}
	}
	return false
}

func intersection(nums1 []int, nums2 []int) []int {
	var is []int
	for _, item1 := range nums1 {
		for _, item2 := range nums2 {
			if item2 == item1 {
				if !hasItem(is, item2) {
					is = append(is, item2)
				}
			}
		}
	}
	return is
}

// @lc code=end

