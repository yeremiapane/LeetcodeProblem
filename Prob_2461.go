package main

import (
	"fmt"
	_ "slices"
)

func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	maxSum := int64(0)
	currentSum := 0
	windowElements := make(map[int]bool)
	left := 0

	for right := 0; right < n; right++ {
		for right-left+1 > k || windowElements[nums[right]] {
			delete(windowElements, nums[left])
			currentSum -= nums[left]
			left++
		}
		windowElements[nums[right]] = true
		currentSum += nums[right]

		if right-left+1 == k {
			if int64(currentSum) > maxSum {
				maxSum = int64(currentSum)
			}
		}
	}
	return maxSum
}

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected int64
	}{
		{nums: []int{1, 5, 4, 2, 9, 9, 9}, k: 3, expected: 15},
		{nums: []int{4, 4, 4}, k: 3, expected: 0},
	}

	for _, tc := range testCases {
		fmt.Printf("Input: nums = %v, k = %d\n", tc.nums, tc.k)
		result := maximumSubarraySum(tc.nums, tc.k)
		fmt.Printf("Output: %v\n", result)
		fmt.Printf("Expected: %v\n\n", tc.expected)
	}
}

/*

func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	if k > n {
		return 0
	}

	maxSum := int64(0)
	for i := 0; i <= n-k; i++ {
		temp := []int{}
		isUnique := true
		elementSet := make(map[int]bool)

		for j := i; j < i+k; j++ {
			if elementSet[nums[j]] {
				isUnique = false
				break
			}
			elementSet[nums[j]] = true
			temp = append(temp, nums[j])
		}

		if isUnique {
			sum := 0
			for _, val := range temp {
				sum += val
			}
			if int64(sum) > maxSum {
				maxSum = int64(sum)
			}
		}
	}

	return maxSum
}
*/
