/*
	Problem : 2461. Maximum Sum of Distinct Subarrays With Length K
	Link 	: https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/
*/

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
	Pseudocode :


	function maximumSubarraySum(nums, k):
    n = length(nums)
    max_sum = 0
    current_sum = 0
    left = 0
    window_elements = empty set

    for right from 0 to n-1:
        while right - left + 1 > k or nums[right] in window_elements:
            remove nums[left] from window_elements
            current_sum -= nums[left]
            left += 1

        add nums[right] to window_elements
        current_sum += nums[right]

        if right - left + 1 == k:
            max_sum = max(max_sum, current_sum)

    return max_sum

function max(a, b):
    if a > b:
        return a
    return b

*/

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
