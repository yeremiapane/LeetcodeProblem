/*
	Problem : 2064. Minimized Maximum of Products Distributed to Any Store
	Link 	: https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/description/?envType=daily-question&envId=2024-11-14
*/

package main

import (
	"fmt"
)

func minimizeMaxQuantity(n int, quantities []int) int {
	left, right := 1, 0
	for _, qty := range quantities {
		if qty > right {
			right = qty
		}
	}

	isPossible := func(mid int) bool {
		storesNeeded := 0
		for _, qty := range quantities {
			storesNeeded += (qty + mid - 1) / mid
		}
		return storesNeeded <= n
	}

	for left < right {
		mid := (left + right) / 2
		if isPossible(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	n := 7
	quantities := []int{15, 10, 10}
	result := minimizeMaxQuantity(n, quantities)
	fmt.Println(result)
}
