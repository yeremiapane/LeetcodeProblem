package main

import (
	"fmt"
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	// Sort items based on price
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	// Compute prefix maximum array
	n := len(items)
	prefixMax := make([]int, n)
	prefixMax[0] = items[0][1]
	for i := 1; i < n; i++ {
		if items[i][1] > prefixMax[i-1] {
			prefixMax[i] = items[i][1]
		} else {
			prefixMax[i] = prefixMax[i-1]
		}
	}

	// Find maximum beauty for each query using binary search
	results := make([]int, len(queries))
	for i, query := range queries {
		// Use binary search to find the rightmost item with price <= query
		idx := sort.Search(n, func(i int) bool {
			return items[i][0] > query
		})
		if idx == 0 {
			results[i] = 0
		} else {
			results[i] = prefixMax[idx-1]
		}
	}

	return results
}

func main() {
	items := [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}
	queries := []int{1, 2, 3, 4, 5}
	fmt.Println(maximumBeauty(items, queries))
}
