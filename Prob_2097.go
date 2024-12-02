package main

import (
	"fmt"
)

func validArrangement(pairs [][]int) [][]int {
	graph := make(map[int][]int)
	edgeIndex := make(map[int]int)

	for _, pair := range pairs {
		start, end := pair[0], pair[1]
		graph[start] = append(graph[start], end)
	}

	outDegree := make(map[int]int)
	inDegree := make(map[int]int)
	for _, pair := range pairs {
		outDegree[pair[0]]++
		inDegree[pair[1]]++
	}
	var startNode int
	for node := range graph {
		if outDegree[node] > inDegree[node] {
			startNode = node
			break
		}
	}
	if startNode == 0 {
		for node := range graph {
			startNode = node
			break
		}
	}

	var stack []int
	stack = append(stack, startNode)
	var result [][]int

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if edgeIndex[node] < len(graph[node]) {
			next := graph[node][edgeIndex[node]]
			edgeIndex[node]++
			stack = append(stack, next)
		} else {

			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				prev := stack[len(stack)-1]
				result = append(result, []int{prev, node})
			}
		}
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}

	return result
}

func main() {
	testCases := []struct {
		pairs    [][]int
		expected [][]int
	}{
		{pairs: [][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}, expected: [][]int{{11, 9}, {9, 4}, {4, 5}, {5, 1}}},
		{pairs: [][]int{{1, 3}, {3, 2}, {2, 1}}, expected: [][]int{{1, 3}, {3, 2}, {2, 1}}},
		{pairs: [][]int{{1, 2}, {1, 3}, {2, 1}}, expected: [][]int{{1, 2}, {2, 1}, {1, 3}}},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: %v\n", testCase.pairs)
		fmt.Printf("Output: %v\n", validArrangement(testCase.pairs))
		fmt.Printf("Expected: %v\n\n", testCase.expected)
	}
}
