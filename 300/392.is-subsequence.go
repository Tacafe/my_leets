/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	if s == "" { return true }
	if t == "" { return false }

	sSlice := strings.Split(s, "")
	tSlice := strings.Split(t, "")
	for i := 0; i < len(tSlice); i++ {
		if len(sSlice) == 0 { return true }

		if sSlice[0] == tSlice[i] {
			sSlice = sSlice[1:]
		}
	}

	return len(sSlice) == 0
}
// @lc code=end

