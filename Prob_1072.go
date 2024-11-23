/*
	Problem : 1072. Flip Columns For Maximum Number of Equal Rows
	Link 	: https://leetcode.com/problems/flip-columns-for-maximum-number-of-equal-rows/description/?envType=daily-question&envId=2024-11-22
*/

package main

import (
	"fmt"
	"strings"
)

func maxEqualRowsAfterFlips(matrix [][]int) int {
	flipsRow := map[string]int{}
	for _, row := range matrix {
		if len(row) == 0 {
			continue
		}
		target := row[0]
		flip := []string{}
		for i, col := range row {
			if col != target {
				flip = append(flip, fmt.Sprintf("%d", i))

			}
		}
		flipsRow[strings.Join(flip, ",")] += 1

	}
	m := 0
	for _, flips := range flipsRow {
		if m < flips {
			m = flips
		}
	}
	return m
}

func main() {
	testCase := []struct {
		matrix   [][]int
		expected int
	}{
		{matrix: [][]int{{0, 1}, {1, 1}}, expected: 1},
		{matrix: [][]int{{0, 1}, {1, 0}}, expected: 2},
		{matrix: [][]int{{0, 0}, {0, 0, 1}, {1, 1, 0}}, expected: 2},
	}

	for _, test := range testCase {
		fmt.Printf("Input Matrix : %v\n", test.matrix)
		fmt.Printf("Output : %d\n", maxEqualRowsAfterFlips(test.matrix))
		fmt.Printf("Expected : %d\n", test.expected)
	}

}
