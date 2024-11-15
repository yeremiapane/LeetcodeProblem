/*
	Problem : 1574. Shortest Subarray to be Removed to Make Array Sorted
	Link 	: https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted/description/?envType=daily-question&envId=2024-11-15
*/

package main

import (
	"fmt"
	"math"
)

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	left := 0

	leftBorder := func() int {
		for left < n-1 && arr[left] <= arr[left+1] {
			left++
		}
		return left
	}

	right := n - 1

	rightBorder := func() int {
		for right > 0 && arr[right] >= arr[right-1] {
			right--
		}
		return right
	}

	left = leftBorder()
	right = rightBorder()

	if left >= right {
		return 0
	}

	result := math.Min(float64(n-left-1), float64(right))

	i, j := 0, right
	for i <= left && j < n {
		if arr[i] <= arr[j] {
			result = math.Min(result, float64(j-i-1))
			i++
		} else {
			j++
		}
	}

	return int(result)
}

func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5})) // Output: 3
	fmt.Println(findLengthOfShortestSubarray([]int{5, 4, 3, 2, 1}))           // Output: 4
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3}))                 // Output: 0
}
