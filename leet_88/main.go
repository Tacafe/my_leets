package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	var newArr []int
	i := 0
	j := 0
	for i < m+n {
		if len(newArr) >= m+n {
			break
		}

		num1 := nums1[i]

		if j >= n {
			newArr = append(newArr, num1)
			i += 1
			continue
		}

		num2 := nums2[j]
		if num1 == 0 {
			newArr = append(newArr, num2)
			i += 1
			j += 1
			continue
		}

		if num1 > num2 {
			newArr = append(newArr, num2)
			j += 1
		} else {
			newArr = append(newArr, num1)
			i += 1
		}
	}
	nums1 = newArr
	fmt.Println(nums1)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
