package main

import (
	"fmt"
	"math"
)

func decrypt(code []int, k int) []int {
	n := len(code)
	result := make([]int, n)

	if k == 0 {
		return result
	}

	for i := 0; i < n; i++ {
		sum := 0
		if k > 0 {
			for j := 1; j <= k; j++ {
				sum += code[(i+j)%n]
			}
		} else {
			for j := 1; j <= int(math.Abs(float64(k))); j++ {
				sum += code[(i-j+n)%n]
			}
		}
		result[i] = sum
	}

	return result
}

func main() {
	testCases := []struct {
		code     []int
		k        int
		expected []int
	}{
		{code: []int{5, 7, 1, 4}, k: 3, expected: []int{12, 10, 16, 13}},
		{code: []int{1, 2, 3, 4}, k: 0, expected: []int{0, 0, 0, 0}},
		{code: []int{2, 4, 9, 3}, k: -2, expected: []int{12, 5, 6, 13}},
	}

	for _, tc := range testCases {
		fmt.Printf("Input: code = %v, k = %d\n", tc.code, tc.k)
		result := decrypt(tc.code, tc.k)
		fmt.Printf("Output: %v\n", result)
		fmt.Printf("Expected: %v\n\n", tc.expected)
	}
}
