/*
	Problem : 1861. Rotating The Box
	Link 	: https://leetcode.com/problems/rotating-the-box/description/?envType=daily-question&envId=2024-11-23
*/

package main

import "fmt"

func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])

	for i := 0; i < m; i++ {
		temp := -1
		for j := n - 1; j >= 0; j-- {
			if box[i][j] == '*' {
				temp = -1
			} else if box[i][j] == '.' && temp == -1 {
				temp = j
			} else if box[i][j] == '#' && temp != -1 {
				box[i][j], box[i][temp] = box[i][temp], box[i][j]
				temp--
			}
		}
	}
	result := make([][]byte, n)
	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			result[i] = append(result[i], box[j][i])
		}
	}
	return result
}

func main() {
	testCases := []struct {
		box      [][]byte
		expected [][]byte
	}{
		{box: [][]byte{{'#', '.', '#'}}, expected: [][]byte{{'.'}, {'#'}, {'#'}}},
		{box: [][]byte{{'#', '.', '*', '.'}, {'#', '#', '*', '.'}}, expected: [][]byte{{'#', '.'}, {'#', '#'}, {'*', '*'}, {'.', '.'}}},
	}
	for _, tC := range testCases {
		fmt.Printf("Input Box : %v\n", tC.box)
		fmt.Printf("Output : %v\n", rotateTheBox(tC.box))
		fmt.Printf("Expected : %v\n", tC.expected)
	}
}
