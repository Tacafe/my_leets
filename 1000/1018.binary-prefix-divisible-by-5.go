package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=1018 lang=golang
 *
 * [1018] Binary Prefix Divisible By 5
 */

// @lc code=start
func prefixesDivBy5(nums []int) []bool {
	ret := []bool{}
	mod := 0

	for i := 0; i < len(nums); i++ {
		mod = (mod*2 + nums[i]) % 5 // prevent overflow
		ret = append(ret, mod == 0)
	}
	return ret
}

// @lc code=end

func main() {
	// test := []int{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1}
	test := []int{1, 0, 1, 0}
	fmt.Println(prefixesDivBy5(test))
}
