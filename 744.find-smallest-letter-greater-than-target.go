/*
 * @lc app=leetcode id=744 lang=golang
 *
 * [744] Find Smallest Letter Greater Than Target
 */

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] <= target {
		return letters[0]
	}

	i := 0
	for i < len(letters)-1 {
		if letters[i] > target {
			break
		}
		i++
	}
	return letters[i]
}

// @lc code=end

