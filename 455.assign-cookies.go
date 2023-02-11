/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	if len(s) == 0 { return 0 }

	sort.Sort(sort.IntSlice(g))
	sort.Sort(sort.IntSlice(s))
	contentCount := 0

	for i := 0; i < len(s); i++ {
		if len(g) == 0 { break }

		if s[i] >= g[0] {
			g = g[1:]
			contentCount += 1
		}
	}
	return contentCount
}
// @lc code=end

