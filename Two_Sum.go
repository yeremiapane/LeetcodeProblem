package main

import "fmt"

func twoSum(nums []int, target int) []int {
	nilai_i := 0
	nilai_j := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			total := nums[i] + nums[j]
			if total == target {
				nilai_i += i
				nilai_j += j
			}
		}
	}
	return []int{nilai_i, nilai_j}
}

func main() {
	test1 := twoSum
	test2 := twoSum
	test3 := twoSum

	fmt.Println(test1([]int{2, 7, 11, 15}, 9))
	fmt.Println(test2([]int{3, 2, 4}, 6))
	fmt.Println(test3([]int{3, 3}, 6))
}
