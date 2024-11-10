/*
	Problem : 3097. Shortest Subarray With OR at Least K II
	Link 	: https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-ii/description/?envType=daily-question&envId=2024-11-10
*/

package main

import (
	"fmt"
	"math"
)

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	bitCount := make([]int, 32)
	currentOR := 0
	left := 0
	minLength := math.Inf(1)

	for right := 0; right < n; right++ {

		currentOR |= nums[right]

		for bit := 0; bit < 32; bit++ {
			if nums[right]&(1<<bit) > 0 {
				bitCount[bit]++
			}
		}

		for left <= right && currentOR >= k {
			minLength = math.Min(minLength, float64(right-left+1))

			updatedOR := 0
			for bit := 0; bit < 32; bit++ {
				if nums[left]&(1<<bit) > 0 {
					bitCount[bit]--
				}
				if bitCount[bit] > 0 {
					updatedOR |= (1 << bit)
				}
			}

			currentOR = updatedOR
			left++
		}
	}

	if minLength == math.Inf(1) {
		return -1
	}
	return int(minLength)
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(minimumSubarrayLength(nums, 3)) // Output: 1 (subarray [3] memenuhi syarat)

	nums = []int{2, 1, 8}
	fmt.Println(minimumSubarrayLength(nums, 10)) // Output: 3 (subarray [2, 1, 8] memenuhi syarat)

	nums = []int{5, 1, 2, 3, 1, 2}
	fmt.Println(minimumSubarrayLength(nums, 7)) // Output: 3 (subarray [5, 1, 2] memenuhi syarat)

	nums = []int{1, 2}
	fmt.Println(minimumSubarrayLength(nums, 0)) // Output: 1 (subarray [1] atau [2] memenuhi syarat)
}
