/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 { return true }

	dp := [len(s)][len(p)]bool{}
	dp[0][0] = true

	for i,  := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
		}
	}

}

// @lc code=end

