/*
	Problem : 3243. Shortest Distance After Road Addition Queries I
	Link 	: https://leetcode.com/problems/shortest-distance-after-road-addition-queries-i/description/?envType=daily-question&envId=2024-11-27
*/

package main

import "fmt"

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	distances := make([]int, n)
	for i := range distances {
		distances[i] = n - 1 - i
	}

	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i+1] = append(graph[i+1], i)
	}
	result := make([]int, 0, len(queries))

	for _, query := range queries {
		source, target := query[0], query[1]

		graph[target] = append(graph[target], source)
		if distances[target]+1 < distances[source] {
			distances[source] = distances[target] + 1
		}

		updateDistances(graph, source, distances)
		result = append(result, distances[0])
	}
	return result
}

func updateDistances(graph [][]int, current int, distances []int) {
	newDist := distances[current] + 1

	for _, nodes := range graph[current] {
		if distances[nodes] <= newDist {
			continue
		}
		distances[nodes] = newDist
		updateDistances(graph, nodes, distances)
	}
}

func main() {
	testCases := []struct {
		n        int
		queries  [][]int
		expected []int
	}{
		{n: 5, queries: [][]int{{2, 4}, {0, 2}, {0, 4}}, expected: []int{3, 2, 1}},
		{n: 4, queries: [][]int{{0, 3}, {0, 2}}, expected: []int{1, 1}},
	}

	for _, tc := range testCases {
		fmt.Printf("Input n : %d\n", tc.n)
		fmt.Printf("Input queries: %v\n", tc.queries)
		fmt.Printf("Output : %v\n", shortestDistanceAfterQueries(tc.n, tc.queries))
		fmt.Printf("Expected: %v\n", tc.expected)
	}
}
