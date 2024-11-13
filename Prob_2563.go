package main

import (
	"fmt"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	n := len(nums)
	count := int64(0)

	for i := 0; i < n; i++ {
		left := lowerBound(nums, lower-nums[i])
		right := upperBound(nums, upper-nums[i])
		count += int64(right - left)
		if lower <= 2*nums[i] && 2*nums[i] <= upper {
			count--
		}
	}
	return count / 2
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func upperBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	nums1 := []int{0, 1, 7, 4, 4, 5}
	lower1, upper1 := 3, 6
	fmt.Println(countFairPairs(nums1, lower1, upper1)) // Output: 6

	nums2 := []int{1, 7, 9, 2, 5}
	lower2, upper2 := 11, 11
	fmt.Println(countFairPairs(nums2, lower2, upper2)) // Output: 1
}

/*
	Ini adalah Code yang benar dalam menghitung, tetapi kompleksitasnya adalah O(n^2) yang
	menyebabkan TLE (Time Limit Exceeded).

	func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	temp := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]+nums[j] >= lower && nums[i]+nums[j] <= upper {
				temp += 1
			}
		}
	}
	return int64(temp)
}
*/
