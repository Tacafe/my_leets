package main

import "fmt"

type Stack []rune

func (stack *Stack) push(c rune) {
	if len(*stack) == 0 {
		*stack = append(*stack, c)
		return
	}

	current := (*stack)[len(*stack)-1]
	if isValidPair(current, c) {
		*stack = append(*stack, c)
		(*stack).pop_pair()
	} else {
		*stack = append(*stack, c)
	}
}

func (stack *Stack) pop_pair() []rune {
	pair := (*stack)[len(*stack)-2:]
	*stack = (*stack)[:len(*stack)-2]
	return pair
}

func isValidPair(current rune, target rune) bool {
	var valid_pair rune

	switch current {
	case '(':
		valid_pair = ')'
	case '{':
		valid_pair = '}'
	case '[':
		valid_pair = ']'
	default:
		valid_pair = '_'
	}
	if valid_pair == '_' {
		return false
	}
	return target == valid_pair
}

func isValid(s string) bool {
	var stack Stack = make([]rune, 0)
	for _, c := range s {
		stack.push(c)
		fmt.Println(stack)
	}
	return len(stack) == 0
}

func main() {
	s := "()(}()"
	fmt.Println(isValid(s))
}
