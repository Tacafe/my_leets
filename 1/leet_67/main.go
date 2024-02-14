package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	var add_by string
	var add_to string
	if len(a) >= len(b) {
		add_by = a
		add_to = b
	} else {
		add_by = b
		add_to = a
	}
	fmt.Println(add_to)

	carry := 0
	answer := ""
	for i := 0; i < len(add_by); i++ {
		by, _ := strconv.Atoi(string(add_by[len(add_by)-1-i]))
		to := 0
		if i < len(add_to) {
			to, _ = strconv.Atoi(string(add_to[len(add_to)-1-i]))
		}

		var d string
		switch by + to + carry {
		case 0:
			carry = 0
			d = "0"
		case 1:
			carry = 0
			d = "1"
		case 2:
			carry = 1
			d = "0"
		case 3:
			carry = 1
			d = "1"
		}
		answer = d + answer
	}

	if carry == 1 {
		answer = "1" + answer
	}
	return answer
}

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}
