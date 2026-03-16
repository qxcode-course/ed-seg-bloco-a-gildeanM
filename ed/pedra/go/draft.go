package main

import (
	"fmt"
)

const MIN_DISTANCE = 10

func calc_distance(stoneA, stoneB int) int {
	diff := stoneA - stoneB
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	rank := make([][2]int, n)
	minor_distance := 101
	idx_winner := -1

	for i := 0; i < n; i++ {
		fmt.Scan(&rank[i][0], &rank[i][1])
	}

	for i := 0; i < n; i++ {
		distance := calc_distance(rank[i][0], rank[i][1])

		if rank[i][0] < MIN_DISTANCE || rank[i][1] < MIN_DISTANCE {
			continue
		}
		if distance < minor_distance {
			minor_distance = distance
			idx_winner = i
		}
	}

	if idx_winner == -1 {
		fmt.Println("sem ganhador")
		return
	}

	fmt.Println(idx_winner)
}
