/*
	Problem : 796. Rotate String
	Link 	: https://leetcode.com/problems/rotate-string/submissions/1453589121/?envType=daily-question&envId=2024-11-14
*/

package main

import (
	"fmt"
	"strings"
)

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	doubled := s + s
	return strings.Contains(doubled, goal)
}

func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
	fmt.Println(rotateString("abcde", "abced"))
}
