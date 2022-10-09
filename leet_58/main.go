package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	vocabs := strings.Split(s, " ")

	i := len(vocabs) - 1
	for {
		if vocabs[i] != "" {
			return len(vocabs[i])
		}
		i--
	}

	return len(vocabs[0])
}

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}
