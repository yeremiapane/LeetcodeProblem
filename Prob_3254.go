/*
	Problem : 3254. Find the Power of K-Size Subarrays I
	Link 	: https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/description/?envType=daily-question&envId=2024-11-16
*/

package main

import (
	"fmt"
	"slices"
	"sort"
)

func isConsecutive(nums []int) bool {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i]+1 != nums[i+1] {
			return false
		}
	}
	return true
}

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}

	temp := []int{}
	for i := 0; i < n-k+1; i++ {
		newArr := []int{}
		for j := 0; j < k; j++ {
			newArr = append(newArr, nums[i+j])
		}
		if k == 1 || (sort.IntsAreSorted(newArr) && isConsecutive(newArr)) {
			temp = append(temp, slices.Max(newArr))
		} else {
			temp = append(temp, -1)
		}
	}
	return temp
}

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{nums: []int{2, 2, 2, 2, 2}, k: 4, expected: []int{-1, -1}},
		{nums: []int{1, 3, 4}, k: 2, expected: []int{-1, 4}},
		{nums: []int{1, 4}, k: 1, expected: []int{1, 4}},
		{nums: []int{1, 2, 3, 21}, k: 3, expected: []int{3, -1}},
		{nums: []int{1, 30, 31, 32}, k: 3, expected: []int{-1, 32}},
	}

	for _, tc := range testCases {
		fmt.Printf("Input: nums = %v, k = %d\n", tc.nums, tc.k)
		result := resultsArray(tc.nums, tc.k)
		fmt.Printf("Output: %v\n", result)
		fmt.Printf("Is Consecutive : %v\n", isConsecutive(tc.nums))
		fmt.Printf("Expected: %v\n\n", tc.expected)
	}
}
