/*
	Problem : Number Complement
	Link : https://leetcode.com/problems/number-complement/submissions/1442857273/?envType=daily-question&envId=2024-11-04
*/

package main

import (
	"fmt"
	"strconv"
)

func findComplement(num int) int {
	mapping := map[rune]rune{'0': '1', '1': '0'}
	biner := strconv.FormatInt(int64(num), 2)

	var value string
	for _, digit := range biner {
		value += string(mapping[digit])
	}

	binars, _ := strconv.ParseInt(value, 2, 64)
	return int(binars)
}

func main() {
	fmt.Println(findComplement(5))
}
