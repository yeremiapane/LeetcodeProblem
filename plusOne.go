package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1] += 1

	for i := n - 1; i > 0; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			digits[i-1] += 1
		}
	}

	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
}
