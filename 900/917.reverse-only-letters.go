/*
 * @lc app=leetcode id=917 lang=golang
 *
 * [917] Reverse Only Letters
 */

// @lc code=start
func isLetter(c int) bool {
	// Upper leter
	if c >= 65 && c <= 90 { return true }
	// Lower leter
	if c >= 97 && c <= 122 { return true }
	return false
}

func reverseOnlyLetters(s string) string {
	if len(s) == 1 { return s }

	letters := []rune(s)

	i := 0
	j := len(letters) - 1
	for i + 1 <= j {
		iIsLetter := isLetter(int(letters[i]))
		jIsLetter := isLetter(int(letters[j]))

		if iIsLetter && jIsLetter {
			il := letters[i]
			jl := letters[j]
			letters[i] = jl
			letters[j] = il
			i += 1
			j -= 1
		} else {
			if !iIsLetter { i += 1 }
			if !jIsLetter { j -= 1 }
		}
	}

	return string(letters)
}
// @lc code=end

