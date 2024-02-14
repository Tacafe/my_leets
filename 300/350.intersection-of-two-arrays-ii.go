/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	var is []int
	for _, item1 := range nums1 {
		for j, item2 := range nums2 {
			if item2 == item1 {
				is = append(is, item2)
				if j < len(nums2)-1 {
					nums2 = append(nums2[:j], nums2[j+1:]...)
				} else if j == len(nums2)-1 {
					nums2 = nums2[:j]
				} else if len(nums2) == 1 {
					nums2 = []int{}
				}
				break
			}
		}
	}
	return is
}

// @lc code=end

