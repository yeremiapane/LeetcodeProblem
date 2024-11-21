package main

import "fmt"

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	var result = int(m * n)
	puzzle := make([][]byte, m)
	for i := 0; i < m; i++ {
		puzzle[i] = make([]byte, n)
	}

	for _, guard := range guards {
		puzzle[guard[0]][guard[1]] = 'G'
		result--
	}

	for _, wall := range walls {
		puzzle[wall[0]][wall[1]] = 'W'
		result--
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if puzzle[i][j] == 'G' {
				for k := i - 1; k >= 0; k-- {
					if puzzle[k][j] == 'R' {
						continue
					} else if puzzle[k][j] != 0 {
						break
					} else {
						puzzle[k][j] = 'R'
						result--
					}
				}
				for k := j + 1; k < m; k++ {
					if puzzle[k][j] == 'R' {
						continue
					} else if puzzle[k][j] != 0 {
						break
					} else {
						puzzle[k][j] = 'R'
						result--
					}
				}
				for k := j - 1; k >= 0; k-- {

					if puzzle[i][k] == 'R' {
						continue
					} else if puzzle[i][k] != 0 {
						break
					} else {
						puzzle[i][k] = 'R'
						result--
					}
				}
				for k := j + 1; k < n; k++ {
					if puzzle[i][k] == 'R' {
						continue
					} else if puzzle[i][k] != 0 {
						break
					} else {
						puzzle[i][k] = 'R'
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
