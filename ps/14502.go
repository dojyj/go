package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var res = 0
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, 1, -1}
var N, M int

type Location struct {
	x int
	y int
}

type Queue struct {
	loc chan Location
}

func (q *Queue) Size() int {
	return len(q.loc)
}

func (q *Queue) Append(i Location) {
	q.loc <- i
}

func (q *Queue) Pop() Location {
	return <-q.loc
}

func bfs(Map [][]int) {
	q := Queue{loc: make(chan Location, 65)}
	newMap := make([][]int, N)
	for idx := range Map {
		newMap[idx] = make([]int, M)
		copy(newMap[idx], Map[idx])
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if newMap[i][j] == 2 {
				q.Append(Location{x: i, y: j})
			}
		}
	}

	// Spread virus
	for q.Size() > 0 {
		size := q.Size()

		for ; size > 0; size-- {
			loc := q.Pop()

			for i := 0; i < 4; i++ {
				next_x, next_y := loc.x+dx[i], loc.y+dy[i]
				if next_x >= 0 && next_y >= 0 && next_x < N && next_y < M && newMap[next_x][next_y] == 0 {
					newMap[next_x][next_y] = 2
					q.Append(Location{x: next_x, y: next_y})
				}
			}
		}
	}

	// find safe place
	cnt := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if newMap[i][j] == 0 {
				cnt++
			}
		}
	}

	if cnt > res {
		res = cnt
	}
}

// Make 3 walls
func makeWall(Map [][]int, cnt int) {
	if cnt == 3 {
		bfs(Map)
		return
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if Map[i][j] == 0 {
				Map[i][j] = 1
				makeWall(Map, cnt+1)
				Map[i][j] = 0
			}
		}
	}
}

func main() {

	fmt.Scanf("%d %d\n", &N, &M)

	Map := make([][]int, N)
	for i := 0; i < N; i++ {
		Map[i] = make([]int, M)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			scanner.Scan()
			Map[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			newMap := make([][]int, N)
			for idx := range Map {
				newMap[idx] = make([]int, M)
				copy(newMap[idx], Map[idx])
			}
			if newMap[i][j] == 0 {
				newMap[i][j] = 1
				makeWall(newMap, 1)
				newMap[i][j] = 0
			}
		}
	}

	fmt.Println(res)
}
