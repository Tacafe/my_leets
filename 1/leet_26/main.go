package main

import "fmt"

func removeDuplicates(nums []int) int {
	i := 0
	for {
		if i >= len(nums)-1 {
			break
		}
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i += 1
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 1, 1}
	fmt.Println(removeDuplicates(nums))
}
