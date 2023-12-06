/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	romanIntMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0
	prevVal := 10000
	for _, letter := range(strings.Split(s, "")) {
		currentVal := romanIntMap[letter]
		if currentVal > prevVal {
			sum -= prevVal
			sum -= prevVal
		}
		sum += currentVal
		prevVal = currentVal
	}

	return sum
}
// @lc code=end

