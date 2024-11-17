/*
	Problem : 862. Shortest Subarray with Sum at Least K
	Link 	: https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/description/?envType=daily-question&envId=2024-11-17
*/

package main

import (
	"fmt"
	"math"
)

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}

	deque := []int{}
	minLength := math.MaxInt32

	for i := 0; i < len(sum); i++ {
		for len(deque) > 0 && sum[i]-sum[deque[0]] >= k {
			minLength = min(minLength, i-deque[0])
			deque = deque[1:]
		}

		// Jaga deque agar tetap terurut
		for len(deque) > 0 && sum[i] <= sum[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
	}

	if minLength == math.MaxInt32 {
		return -1
	}
	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{nums: []int{1}, k: 1, expected: 1},
		{nums: []int{1, 2}, k: 4, expected: -1},
		{nums: []int{2, -1, 2}, k: 3, expected: 3},
	}

	for _, tc := range testCases {
		fmt.Printf("Input: nums = %v, k = %d\n", tc.nums, tc.k)
		result := shortestSubarray(tc.nums, tc.k)
		fmt.Printf("Output: %v\n", result)
		fmt.Printf("Expected: %v\n\n", tc.expected)
	}
}

/*
	Algoritma :
	function shortestSubarray(nums, k):
    n = length(nums)
    sum = array of size (n + 1) initialized to 0
    for i from 0 to n-1:
        sum[i+1] = sum[i] + nums[i]

    deque = empty array
    minLength = infinity

    for i from 0 to length(sum) - 1:
        while deque is not empty and sum[i] - sum[deque[0]] >= k:
            minLength = min(minLength, i - deque[0])
            remove the first element from deque

        while deque is not empty and sum[i] <= sum[deque[last]]:
            remove the last element from deque

        append i to deque

    if minLength == infinity:
        return -1
    return minLength

function min(a, b):
    if a < b:
        return a
    return b
*/
