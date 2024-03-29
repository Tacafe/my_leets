/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	i := m - 1
	j := n - 1
	k := (m + n) - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--; k--
		} else {
			nums1[k] = nums2[j]
			j--; k--
		}
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--; k--
	}
}

// @lc code=end

