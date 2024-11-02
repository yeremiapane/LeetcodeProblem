package main

import "fmt"

func plusOne(digits []int) []int {
	total := []int{}
	for i := 0; i < len(digits); i++ {
		total = append(total, i)
	}
	return total
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
}
