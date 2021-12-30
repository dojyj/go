package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var N, M, res int

// 0,		1,		2,		3
// north, 	east, 	south, 	west
var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

type Robot struct {
	x int
	y int
	d int
}

func clean(r Robot, place [][]int) {
	for {
		this_x := r.x
		this_y := r.y
		cleaned := false

		// Robot Move Debugger
		// time.Sleep(time.Second)
		// for i := range place {
		// 	fmt.Println(place[i], r.d)
		// }
		// fmt.Println("")

		// #1 clean this place (dirty)
		if place[this_x][this_y] == 0 {
			res++
			place[this_x][this_y] = 2
		}

		// #2.a, #2.b if left of direction is dirty, rotate and clean
		for i := 0; i < 4; i++ {
			left := (4 + r.d - 1) % 4
			left_x, left_y := this_x+dx[left], this_y+dy[left]

			if left_x >= 0 && left_y >= 0 && left_x < N && left_y < M {
				r.d = left
				if place[left_x][left_y] == 0 {
					r.x = left_x
					r.y = left_y
					cleaned = true
					break
				}
			}
		}

		// #2.c, #2.d Backwoard Strategy
		if !cleaned {
			back := (4 + r.d - 2) % 4
			back_x, back_y := r.x+dx[back], r.y+dy[back]

			if back_x >= 0 && back_y >= 0 && back_x < N && back_y < M {
				if place[back_x][back_y] == 1 {
					break
				} else {
					r.x = back_x
					r.y = back_y
				}
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Scanf("%d %d\n", &N, &M)

	var r Robot

	fmt.Scanf("%d %d %d\n", &r.x, &r.y, &r.d)

	place := make([][]int, N)
	for i := 0; i < N; i++ {
		place[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			scanner.Scan()
			place[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	clean(r, place)
	fmt.Println(res)
}
