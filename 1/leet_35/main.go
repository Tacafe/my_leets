package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i, _ := range nums {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(searchInsert(nums, target))
}
