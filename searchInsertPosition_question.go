package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}

func main() {
	// Kasus uji 1
	nums1 := []int{1, 3, 5, 6}
	target1 := 5
	fmt.Println(searchInsert(nums1, target1))
}
