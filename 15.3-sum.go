/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	triplets := [][]int{}
	tripletMap := make(map[[3]int][]int)

	negatives, zeroes, positives := []int{}, []int{}, []int{}
	for _, n := range nums {
		if n < 0 {
			negatives = append(negatives, n)
		} else if n == 0 {
			zeroes = append(zeroes, n)
		} else {
			positives = append(positives, n)
		}
	}

	if len(zeroes) >= 3 {
		triplets = append(triplets, []int{0, 0, 0})
	}

	if len(negatives) == 0 || len(positives) == 0 {
		return triplets
	}

	if len(negatives) > 1 {
		for i := 0; i < len(negatives)-1; i++ {
			for j := i + 1; j < len(negatives); j++ {
				for k := 0; k < len(positives); k++ {
					target := [3]int{negatives[i], negatives[j], positives[k]}
					sort.Ints(target[:])
					s := negatives[i] + negatives[j] + positives[k]
					if s == 0 {
						tripletMap[target] = []int{negatives[i], negatives[j], positives[k]}
					}
				}
			}
		}
	}

	if len(positives) > 1 {
		for i := 0; i < len(positives)-1; i++ {
			for j := i + 1; j < len(positives); j++ {
				for k := 0; k < len(negatives); k++ {
					target := [3]int{positives[i], positives[j], negatives[k]}
					sort.Ints(target[:])
					s := positives[i] + positives[j] + negatives[k]
					if s == 0 {
						tripletMap[target] = []int{positives[i], positives[j], negatives[k]}
					}
				}
			}
		}
	}

	if len(zeroes) > 0 {
		usedPos := []int{}
		for i := 0; i < len(negatives); i++ {
			for j := 0; j < len(positives); j++ {
				if negatives[i]+positives[j] == 0 {
					isUsed := false
					for _, p := range usedPos {
						if p == positives[j] {
							isUsed = true
							break
						}
					}
					if !isUsed {
						tripletMap[[3]int{negatives[i], 0, positives[j]}] = []int{negatives[i], 0, positives[j]}
						usedPos = append(usedPos, positives[j])
					}
				}
			}
		}
	}

	for _, v := range tripletMap {
		triplets = append(triplets, v)
	}

	return triplets
}

// @lc code=end

