package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for {
		if i >= len(nums) {
			break
		}
		if nums[i] == val {
			if i == 0 {
				nums = nums[i+1:]
			} else {
				nums = append(nums[:i], nums[i+1:]...)
			}
		} else {
			i++
		}
	}
	return len(nums)
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}
