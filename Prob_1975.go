package main

import (
	"fmt"
	"math"
)

func maxMatrixSum(matrix [][]int) int64 {
	totalSum := 0
	minAbsValues := math.MaxInt64
	negativeCount := 0

	for _, row := range matrix {
		for _, value := range row {
			totalSum += int(math.Abs(float64(value)))
			if value < 0 {
				negativeCount++
			}
			if int(math.Abs(float64(value))) < minAbsValues {
				minAbsValues = int(math.Abs(float64(value)))
			}
		}
	}

	if negativeCount%2 == 0 {
		return int64(totalSum)
	} else {
		return int64(totalSum - 2*minAbsValues)
	}
}

func main() {
	testCase := []struct {
		matrix   [][]int
		expected int64
	}{
		{matrix: [][]int{{1, -1}, {-1, 1}}, expected: 4},
		{matrix: [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}, expected: 16},
	}

	for _, v := range testCase {
		fmt.Printf("Input Matrix : %v\n", v.matrix)
		fmt.Printf("Output : %v\n", maxMatrixSum(v.matrix))
		fmt.Printf("Expected : %v\n", v.expected)
	}
}
