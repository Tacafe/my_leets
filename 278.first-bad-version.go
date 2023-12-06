/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	if isBadVersion(1) { return 1 }
	if isBadVersion(n) { return n }

	searchMax := int(math.Log2(float64(n))) + 1
	rangeFrom := 1
	rangeTo := n

	for i := 0; i < searchMax; i++ {
		check := int((rangeFrom + rangeTo) / 2)
		i += 1
		prevBad := isBadVersion(check)
		nextBad := isBadVersion(check + 1)

		if !prevBad && nextBad {
			return check + 1
		} else if prevBad && nextBad {
			rangeTo = check - 1
		} else {
			rangeFrom = check + 1
		}
	}
	return n
}
// @lc code=end

