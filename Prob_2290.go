/*
	Problem : 2290. Minimum Obstacle Removal to Reach Corner
	Link 	: https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/description/?envType=daily-question&envId=2024-11-28
*/

package main

type Coord struct {
	i int
	j int
}

func minimumObstacles(grid [][]int) int {
	dist := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		dist[i] = make([]int, len(grid[0]))
		for j := 0; j < len(grid[i]); j++ {
			dist[i][j] = 100000
		}
	}

	q := make([]Coord, 0)
	q = append(q, Coord{i: 0, j: 0})
	dist[0][0] = 0
	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		if s.i-1 >= 0 && dist[s.i][s.j]+grid[s.i-1][s.j] < dist[s.i-1][s.j] {
			dist[s.i-1][s.j] = dist[s.i][s.j] + grid[s.i-1][s.j]
			q = append(q, Coord{i: s.i - 1, j: s.j})
		}
		if s.i+1 < len(grid) && dist[s.i][s.j]+grid[s.i+1][s.j] < dist[s.i+1][s.j] {
			dist[s.i+1][s.j] = dist[s.i][s.j] + grid[s.i+1][s.j]
			q = append(q, Coord{i: s.i + 1, j: s.j})
		}
		if s.j-1 >= 0 && dist[s.i][s.j]+grid[s.i][s.j-1] < dist[s.i][s.j-1] {
			dist[s.i][s.j-1] = dist[s.i][s.j] + grid[s.i][s.j-1]
			q = append(q, Coord{i: s.i, j: s.j - 1})
		}
		if s.j+1 < len(grid[0]) && dist[s.i][s.j]+grid[s.i][s.j+1] < dist[s.i][s.j+1] {
			dist[s.i][s.j+1] = dist[s.i][s.j] + grid[s.i][s.j+1]
			q = append(q, Coord{i: s.i, j: s.j + 1})
		}
	}

	return dist[len(grid)-1][len(grid[0])-1]
}
