package main

import (
	"fmt"
)

func IsPrimeNumber(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeSubOperation(nums []int) bool {
	primes := []int{}
	for i := 2; i < 1000; i++ {
		if IsPrimeNumber(i) {
			primes = append(primes, i)
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			canSubtract := false
			for _, p := range primes {
				if p >= nums[i] {
					break
				}
				if nums[i]-p < nums[i+1] {
					nums[i] -= p
					canSubtract = true
					break
				}
			}
			if !canSubtract {
				return false
			}
		}
	}
	return true

}

func main() {
	nums1 := []int{15, 20, 17, 7, 16}
	fmt.Println(primeSubOperation(nums1)) // Expected output: true

	nums2 := []int{4, 9, 6, 10}
	fmt.Println(primeSubOperation(nums2)) // Expected output: true

	nums3 := []int{6, 8, 11, 12}
	fmt.Println(primeSubOperation(nums3)) // Expected output: true

	nums4 := []int{5, 8, 3}
	fmt.Println(primeSubOperation(nums4)) // Expected output: false
}
