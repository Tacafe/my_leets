package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	carry := true
	i := len(digits) - 1

	for carry {
		if i < 0 {
			digits = append([]int{1}, digits...)
			carry = false
			break
		}

		d := digits[i]
		if d == 9 {
			digits[i] = 0
			i--
		} else {
			digits[i] += 1
			carry = false
		}
	}

	return digits
}

func main() {
	digits := []int{9, 9, 9, 9, 9}
	fmt.Println(plusOne(digits))
}
