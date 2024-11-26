/*
	Problem : 773. Sliding Puzzle
	Link 	: https://leetcode.com/problems/sliding-puzzle/description/?envType=daily-question&envId=2024-11-25
*/

package main

import (
	"fmt"
	"strings"
)

type State struct {
	board          string
	zeroPos, steps int
}

func slidingPuzzle(board [][]int) int {
	target := "123450"
	start := ""
	for _, row := range board {
		for _, cell := range row {
			start += fmt.Sprint(cell)
		}
	}

	neighbors := map[int][]int{
		0: {1, 3},
		1: {0, 2, 4},
		2: {1, 5},
		3: {0, 4},
		4: {1, 3, 5},
		5: {2, 4},
	}

	queue := []State{{start, strings.Index(start, "0"), 0}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.board == target {
			return current.steps
		}

		for _, neighbor := range neighbors[current.zeroPos] {
			newBoard := []rune(current.board)
			newBoard[current.zeroPos], newBoard[neighbor] = newBoard[neighbor], newBoard[current.zeroPos]
			newBoardStr := string(newBoard)

			if !visited[newBoardStr] {
				visited[newBoardStr] = true
				queue = append(queue, State{newBoardStr, neighbor, current.steps + 1})
			}
		}
	}

	return -1
}

func main() {
	testCases := []struct {
		board    [][]int
		expected int
	}{
		{board: [][]int{{1, 2, 3}, {4, 0, 5}}, expected: 1},
		{board: [][]int{{1, 2, 3}, {5, 4, 0}}, expected: -1},
		{board: [][]int{{4, 1, 2}, {5, 0, 3}}, expected: 5},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input Case : %v\n", testCase.board)
		fmt.Printf("Output : %v\n", slidingPuzzle(testCase.board))
		fmt.Printf("Expected : %d\n\n", testCase.expected)
	}
}
