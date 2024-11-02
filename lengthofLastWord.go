package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	spliting := strings.Split(s, " ")
	length := 0
	for i := len(spliting) - 1; i >= 0; i-- {
		if len(spliting[i]) != 0 {
			length = len(spliting[i])
			break
		}
	}
	return length

}

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}
