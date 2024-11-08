/*
	Problem : 1829. Maximum XOR for each query
	link :https://leetcode.com/problems/maximum-xor-for-each-query/submissions/1446855796/?envType=daily-question&envId=2024-11-08
*/

package main

import (
	"fmt"
)

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	result := make([]int, n)
	maxVal := (1 << maximumBit) - 1

	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}

	for i := 0; i < n; i++ {
		result[i] = xorSum ^ maxVal
		xorSum ^= nums[n-1-i]
	}

	return result
}

func main() {
	nums := []int{0, 1, 1, 3}
	maximumBit := 2
	fmt.Println(getMaximumXor(nums, maximumBit)) // Output: [0, 3, 2, 3]
	nums2 := []int{2, 3, 4, 7}
	maximumBit2 := 3
	fmt.Println(getMaximumXor(nums2, maximumBit2)) // Output: [5,2,6,5]
}
