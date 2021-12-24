package main

import (
	"fmt"
	"strings"
)

type location struct {
	x int
	y int
}

var holeLoc location

// 0,1,2,3 => 동,서,남,북
func tilt(board []string, loc int, rloc, bloc location) ([]string, location, location) {

}

func solve(board []string, depth int, rLoc, bLoc location) int {
	ret := -1

	for i := 0; i < 4; i++ {
		tiltedBoard, r, b := tilt(board, i, rLoc, bLoc)
		res := solve(tiltedBoard, depth+1, r, b)
		if res > ret {
			ret = res
		}
	}
	return ret
}

func main() {
	var N, M int
	var rLoc, bLoc location

	fmt.Scanf("%d %d\n", &N, &M)

	board := make([]string, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&board[i])
		if strings.Contains(board[i], "R") {
			rLoc.x = i
			rLoc.y = strings.Index(board[i], "R")
		}

		if strings.Contains(board[i], "B") {
			bLoc.x = i
			bLoc.y = strings.Index(board[i], "B")
		}

		if strings.Contains(board[i], "O") {
			holeLoc.x = i
			holeLoc.y = strings.Index(board[i], "O")
		}
	}

	res := solve(board, 0, rLoc, bLoc)
	fmt.Println(res)
}
