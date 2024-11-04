package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

/*
Or you can use this for the actual remove duplicate

	func removeDuplicates(nums []int) int {
		if len(nums) == 0 {
			return len(nums)
		}

		temp := []int{nums[0]}
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] {
				temp = append(temp, nums[i])
			}
		}
		return len(temp)
	}
*/
func main() {
	//test := []int{1, 1, 2}
	//fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//fmt.Println(removeDuplicates(test))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
