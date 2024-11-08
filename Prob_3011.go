package main

import (
	"fmt"
	"math/bits"
)

func canSortArray(nums []int) bool {
	changed := true

	for changed {
		changed = false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				// Cek apakah elemen yang bersebelahan memiliki jumlah set bits yang sama
				if bits.OnesCount(uint(nums[i])) == bits.OnesCount(uint(nums[i+1])) {
					// Tukar posisi elemen
					nums[i], nums[i+1] = nums[i+1], nums[i]
					changed = true
				} else {
					// Jika tidak, tidak mungkin array bisa diurutkan
					return false
				}
			}
		}
	}
	return true
}

func main() {
	examples := [][]int{
		{8, 4, 2, 30, 15},
		{1, 2, 3, 4, 5},
		{3, 16, 8, 4, 2},
	}

	for _, nums := range examples {
		fmt.Printf("Input: %v, Can be sorted: %v\n", nums, canSortArray(nums))
	}
}
