/*
	Problem : 1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence
	Link 	: https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/description/?envType=daily-question&envId=2024-12-02
*/

package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	nSearchWord := len(searchWord)
	splitString := strings.Split(sentence, " ")
	for index, word := range splitString {
		if len(word) >= nSearchWord {
			if strings.HasPrefix(word, searchWord) {
				return index + 1
			}
		}
	}
	return -1
}

func main() {
	testCase := []struct {
		sentence   string
		searchWord string
		expected   int
	}{
		{sentence: "i love eating burger", searchWord: "burg", expected: 4},
		{sentence: "this problem is an easy problem", searchWord: "pro", expected: 2},
		{sentence: "i am tired", searchWord: "you", expected: -1},
	}

	for _, tc := range testCase {
		fmt.Printf("Input Sentence : %s\n", tc.sentence)
		fmt.Printf("Input Search Word : %s\n", tc.searchWord)
		fmt.Printf("Output : %d\n", isPrefixOfWord(tc.sentence, tc.searchWord))
		fmt.Printf("Expected : %d\n", tc.expected)
	}
}
