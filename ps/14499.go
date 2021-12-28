package main

import (
	"fmt"
)

type Dice struct {
	top, bottom, left, right, front, back int
}

var dx = []int{0, 0, -1, 1}
var dy = []int{1, -1, 0, 0}

func (m *Dice) Turn(direction int, x, y *int, Map [][]int) bool {
	// Exception Handling
	next_x := *x + dx[direction-1]
	next_y := *y + dy[direction-1]
	if next_x < 0 || next_y < 0 || next_x >= len(Map) || next_y >= len(Map[0]) {
		return false
	}

	// If map value == 0, Copy dice bottom value to map
	// Else copy map value to dice bottom, map value becomes 0
	if Map[*x][*y] == 0 {
		Map[*x][*y] = m.bottom
	} else {
		m.bottom = Map[*x][*y]
		Map[*x][*y] = 0
	}

	*x = next_x
	*y = next_y

	// 1, 2, 3, 4 // east, west, north, south
	base := m.bottom
	switch direction {
	case 1:
		m.bottom = m.right
		m.right = m.top
		m.top = m.left
		m.left = base
	case 2:
		m.bottom = m.left
		m.left = m.top
		m.top = m.right
		m.right = base
	case 3:
		m.bottom = m.back
		m.back = m.top
		m.top = m.front
		m.front = base
	case 4:
		m.bottom = m.front
		m.front = m.top
		m.top = m.back
		m.back = base
	}
	return true
}

func main() {
	var N, M, x, y, K int

	fmt.Scanf("%d %d %d %d %d\n", &N, &M, &x, &y, &K)

	Map := make([][]int, N)
	for i := 0; i < N; i++ {
		Map[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Scan(&Map[i][j])
		}
	}

	dice := Dice{}
	for ; K > 0; K-- {
		var cmd int

		fmt.Scan(&cmd)

		if !dice.Turn(cmd, &x, &y, Map) {
			continue
		}

		fmt.Println(dice.top)
	}
}
