package main

import "fmt"

func isPalindrome(x int) bool {
	reverse := 0
	temp := x
	for x > 0 {
		reverse = (reverse * 10) + (x % 10)
		x /= 10
	}
	return temp == reverse
}

func main() {
	fmt.Println(isPalindrome(121))
}
