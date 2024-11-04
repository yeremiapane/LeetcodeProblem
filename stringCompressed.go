/*
	Challenge 1 : Compressed String III
	Link : https://leetcode.com/problems/string-compression-iii/?envType=daily-question&envId=2024-11-04
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compressedString(word string) string {
	split := strings.Split(word, "")
	var n int = len(split)
	var i int = 0
	words := make([]string, n)

	for i < n {
		ch := split[i]
		cnt := 0
		for i < n && split[i] == ch && cnt < 9 {
			cnt++
			i++
		}
		words = append(words, (strconv.Itoa(cnt) + ch))

	}
	return strings.Join(words, "")
}

func main() {
	fmt.Println(compressedString("abcde"))
	fmt.Println(compressedString("aaaaaaaaaaaaaabb"))
}
