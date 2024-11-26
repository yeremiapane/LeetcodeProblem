/*
	Problem 	: 2924. Find Champion II
	Link 		: https://leetcode.com/problems/find-champion-ii/
*/

package main

import "fmt"

func findChampion(n int, edges [][]int) int {
	unDefeated := make([]bool, n)
	for i := range unDefeated {
		unDefeated[i] = true
	}

	for _, edge := range edges {
		_, loser := edge[0], edge[1]
		unDefeated[loser] = false
	}
	champions := -1
	championsCount := 0

	for team := 0; team < n; team++ {
		if unDefeated[team] {
			champions = team
			championsCount++
		}
	}
	if championsCount == 1 {
		return champions
	}
	return -1
}

func main() {
	testCase := []struct {
		n        int
		edges    [][]int
		expected int
	}{
		{n: 3, edges: [][]int{{0, 1}, {1, 2}}, expected: 0},
		{n: 4, edges: [][]int{{0, 2}, {1, 3}, {1, 2}}, expected: -1},
	}

	for _, tc := range testCase {
		fmt.Printf("Input n : %d\n", tc.n)
		fmt.Printf("Input Edges : %v\n", tc.edges)
		fmt.Printf("Output : %d\n", findChampion(tc.n, tc.edges))
		fmt.Printf("Expected : %d\n\n", tc.expected)
	}
}
