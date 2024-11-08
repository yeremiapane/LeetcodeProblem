package main

import (
	"fmt"
	"sort"
)

func countSetBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// Fungsi untuk memeriksa apakah array dapat diurutkan dengan operasi yang valid
func canSortArray(nums []int) bool {
	// Salin array asli untuk dibandingkan nanti
	original := make([]int, len(nums))
	copy(original, nums)

	// Urutkan array berdasarkan jumlah bit yang diset
	sort.Slice(nums, func(i, j int) bool {
		return countSetBits(nums[i]) < countSetBits(nums[j])
	})

	// Periksa apakah array yang diurutkan sama dengan array asli yang diurutkan
	for i := range nums {
		if nums[i] != original[i] {
			return false
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
