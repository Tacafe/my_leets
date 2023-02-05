package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isContain(s []string, target string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == target {
			return i
		}
	}
	return -1
}

func isAnagram(s string, t string) bool {
	sl_s := strings.Split(s, "")
	sl_t := strings.Split(t, "")

	for i := 0; i < len(sl_s); i++ {
		target := sl_s[i]
		pos := isContain(sl_t, target)
		if pos >= 0 {
			sl_t = append(sl_t[:pos], sl_t[pos+1:]...)
		} else {
			return false
		}
	}

	if len(sl_t) == 0 {
		return true
	}
	return false
}

// @lc code=end

func main() {
	str := "abcdefg"
	slice := strings.Split(str, "")

	for i := 0; i < len(slice); i++ {
		fmt.Println(i)
		fmt.Println(slice[i])
	}
}
