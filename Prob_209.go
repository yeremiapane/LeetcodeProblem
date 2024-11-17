/*
	Problem : 209. Minimum Size Subarray Sum
	Link 	: https://leetcode.com/problems/minimum-size-subarray-sum/description/
*/

package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left := 0
	current_sum := 0
	min_length := math.MaxInt32

	for right := 0; right < n; right++ {
		current_sum += nums[right]

		for current_sum >= target {
			min_length = min(min_length, right-left+1)
			current_sum -= nums[left]
			left++
		}
	}

	if min_length == math.MaxInt32 {
		return 0
	}
	return min_length
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	testCases := []struct {
		target   int
		nums     []int
		expected int
	}{
		{target: 7, nums: []int{2, 3, 1, 2, 4, 3}, expected: 2},
		{target: 4, nums: []int{1, 4, 4}, expected: 1},
		{target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}, expected: 0},
	}

	for _, tc := range testCases {
		fmt.Printf("Input: target = %d, nums = %v\n", tc.target, tc.nums)
		result := minSubArrayLen(tc.target, tc.nums)
		fmt.Printf("Output: %v\n", result)
		fmt.Printf("Expected: %v\n\n", tc.expected)
	}
}

/*
	Algoritma :
function minSubArrayLen(target, nums):
    n = length(nums)
    left = 0
    current_sum = 0
    min_length = infinity

    for right from 0 to n-1:
        current_sum += nums[right]

        while current_sum >= target:
            min_length = min(min_length, right - left + 1)
            current_sum -= nums[left]
            left += 1

    if min_length == infinity:
        return 0
    return min_length

function min(a, b):
    if a < b:
        return a
    return b

*/
