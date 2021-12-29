package main

import (
	"fmt"
)

type Apple struct {
	row int
	col int
}

type Dummy struct {
	x int
	y int
}

var cur_time = 1
var cur_direction = 0

// right, down, left, up
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

// remove element at index
func remove(arr []Apple, idx int) []Apple {
	return append(arr[:idx], arr[idx+1:]...)
}

func main() {
	var N, K, L int
	dummy := make([]Dummy, 0)

	fmt.Scan(&N)

	board := make([][]int, N)
	for i := 0; i < N; i++ {
		board[i] = make([]int, N)
	}

	fmt.Scan(&K)

	apple := make([]Apple, K)
	for i := 0; i < K; i++ {
		fmt.Scanf("%d %d\n", &apple[i].row, &apple[i].col)
	}

	fmt.Scan(&L)

	dummy = append(dummy, Dummy{x: 0, y: 0})
	end := false

	for ; L > 0; L-- {
		var t int
		var d string

		fmt.Scanf("%d %s\n", &t, &d)

		if end {
			continue
		}

	EXIT:
		for ; ; cur_time++ {
			dummylen := len(dummy) - 1
			tail_x := dummy[dummylen].x
			tail_y := dummy[dummylen].y

			// fmt.Println("curtime : ", cur_time, "dummy :", dummy)

			// case 1. check if i bump into my body (important! head moves first, tail moves later)
			for i := 1; i <= dummylen; i++ {
				// fmt.Println("DEBUG => ", dummy[i].x, dummy[0].x+dx[cur_direction], dummy[i].y, dummy[0].y+dy[cur_direction])
				if dummy[i].x == dummy[0].x+dx[cur_direction] && dummy[i].y == dummy[0].y+dy[cur_direction] {
					end = true
					break EXIT
				}
			}

			// tail & head moves
			for i := dummylen; i > 0; i-- {
				dummy[i].x, dummy[i].y = dummy[i-1].x, dummy[i-1].y
			}
			dummy[0].x, dummy[0].y = dummy[0].x+dx[cur_direction], dummy[0].y+dy[cur_direction]

			// case 2. check if i bump into the wall
			if dummy[0].x < 0 || dummy[0].x >= N || dummy[0].y < 0 || dummy[0].y >= N {
				end = true
				break EXIT
			}

			// case 3. check if i eat apple. if apple exists, remove apple
			for i, v := range apple {
				if v.row-1 == dummy[0].x && v.col-1 == dummy[0].y {
					dummy = append(dummy, Dummy{x: tail_x, y: tail_y})
					apple = remove(apple, i)
				}
			}

			// until rotate
			if t == cur_time {
				cur_time++
				break EXIT
			}
		}

		switch d {
		case "D":
			cur_direction = (cur_direction + 1) % 4
		case "L":
			cur_direction = (4 + cur_direction - 1) % 4
		}
	}

	// fmt.Println(dummy)
	// fmt.Println(apple)

	// after all rotating
	if !end {
	EXIT2:
		for ; ; cur_time++ {
			dummylen := len(dummy) - 1
			tail_x := dummy[dummylen].x
			tail_y := dummy[dummylen].y

			// fmt.Println("curtime : ", cur_time, "dummy :", dummy)

			for i := 1; i <= dummylen; i++ {
				// fmt.Println("DEBUG => ", dummy[i].x, dummy[0].x+dx[cur_direction], dummy[i].y, dummy[0].y+dy[cur_direction])
				if dummy[i].x == dummy[0].x+dx[cur_direction] && dummy[i].y == dummy[0].y+dy[cur_direction] {
					end = true
					break EXIT2
				}
			}

			for i := dummylen; i > 0; i-- {
				dummy[i].x, dummy[i].y = dummy[i-1].x, dummy[i-1].y
			}

			dummy[0].x, dummy[0].y = dummy[0].x+dx[cur_direction], dummy[0].y+dy[cur_direction]

			if dummy[0].x < 0 || dummy[0].x >= N || dummy[0].y < 0 || dummy[0].y >= N {
				break EXIT2
			}

			for i, v := range apple {
				if v.row-1 == dummy[0].x && v.col-1 == dummy[0].y {
					dummy = append(dummy, Dummy{x: tail_x, y: tail_y})
					apple = remove(apple, i)
				}
			}
		}
	}

	fmt.Println(cur_time)
}
