package main

import "fmt"

func letterToInt(l rune) int {
	switch l {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func romanToInt(s string) int {
	sum := 0
	stack := 0
	l_prev := ' '

	for i, l := range s {
		val = letterToInt(l)

		if (i == 0) {
			stack += val
		}
		stack += letterToInt(l)
		if

	}
	return sum
}

func main() {
	s := "IIIV"
	fmt.Printf("%d\n", romanToInt(s))
}
