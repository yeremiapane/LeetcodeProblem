/*
	Problem : 2257. Count Unguarded Cells in the Grid
	Link 	: https://leetcode.com/problems/count-unguarded-cells-in-the-grid/description/
*/

package main

import "fmt"

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {

	var result int64 = int64(m * n)
	puzzle := make([][]byte, m)
	for i := 0; i < m; i++ {
		puzzle[i] = make([]byte, n)
	}

	for _, guard := range guards {
		puzzle[guard[0]][guard[1]] = 'G'
		result--
	}

	for _, walls := range walls {
		puzzle[walls[0]][walls[1]] = 'W'
		result--
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if puzzle[i][j] == 'G' {
				// up
				for h := i - 1; h >= 0; h-- {
					if puzzle[h][j] == 'R' {
						continue
					} else if puzzle[h][j] != 0 {
						break
					} else {
						puzzle[h][j] = 'R'
						result--
					}
				}

				// down
				for h := i + 1; h < m; h++ {

					if puzzle[h][j] == 'R' {
						continue
					} else if puzzle[h][j] != 0 {
						break
					} else {
						puzzle[h][j] = 'R'
						result--
					}
				}

				// left
				for h := j - 1; h >= 0; h-- {

					if puzzle[i][h] == 'R' {
						continue
					} else if puzzle[i][h] != 0 {
						break
					} else {
						puzzle[i][h] = 'R'
						result--
					}
				}
				// right
				for h := j + 1; h < n; h++ {
					if puzzle[i][h] == 'R' {
						continue
					} else if puzzle[i][h] != 0 {
						break
					} else {
						puzzle[i][h] = 'R'
						result--
					}
				}
			}
		}
	}

	return int(result)
}

func main() {
	testCases := []struct {
		m, n     int
		guards   [][]int
		walls    [][]int
		expected int
	}{
		{m: 4, n: 6, guards: [][]int{{0, 0}, {1, 1}, {2, 3}}, walls: [][]int{{0, 1}, {2, 2}, {1, 4}}, expected: 7},
		{m: 3, n: 3, guards: [][]int{{1, 1}}, walls: [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}, expected: 0},
	}

	for _, tc := range testCases {
		fmt.Printf("Input: m = %v, n = %d\n", tc.m, tc.n)
		result := countUnguarded(tc.m, tc.n, tc.guards, tc.walls)
		fmt.Printf("Output: %v\n", result)
		fmt.Printf("Expected: %v\n\n", tc.expected)
	}
}
