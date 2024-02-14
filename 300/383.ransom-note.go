package main

import "fmt"

func remove(s []rune, r rune) []rune {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func canConstruct(ransomNote string, magazine string) bool {
	rRunes := []rune(ransomNote)
	mRunes := []rune(magazine)
	target := []rune(ransomNote)

	for _, r := range rRunes {
		isContain := false
		for _, m := range mRunes {
			if r == m {
				isContain = true
			}
		}

		if isContain {
			mRunes = remove(mRunes, r)
			target = remove(target, r)
		}
	}
	return len(target) == 0
}

func main() {
	ransomNote := "aab"
	magazine := "baa"
	fmt.Println(canConstruct(ransomNote, magazine))
}
