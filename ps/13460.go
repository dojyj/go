package main

import (
	"fmt"
	"math"
	"strings"
)

type location struct {
	x int
	y int
}

type Item struct {
	Red  location
	Blue location
}

type Queue struct {
	it chan Item
}

func (q *Queue) Size() int {
	return len(q.it)
}

func (q *Queue) Append(i Item) {
	q.it <- i
}

func (q *Queue) Pop() Item {
	return <-q.it
}

var rLoc, bLoc location

// 상, 우, 하, 좌
var dx = [...]int{-1, 0, 1, 0}
var dy = [...]int{0, 1, 0, -1}

func bfs(board []string) int {
	ret := 0
	visited := make(map[string]bool) // if N, M > 10 => visited must be [N][M][N][M] bool
	s := fmt.Sprintf("%d%d%d%d", rLoc.x, rLoc.y, bLoc.x, bLoc.y)
	visited[s] = true
	q := Queue{it: make(chan Item, 100)}
	q.Append(Item{Red: rLoc, Blue: bLoc})

	for q.Size() > 0 {
		size := q.Size()
		for ; size > 0; size-- {
			it := q.Pop()
			rx := it.Red.x
			ry := it.Red.y
			bx := it.Blue.x
			by := it.Blue.y

			// fmt.Println("new queue", rx, ry, bx, by)

			// 구멍을 만났고, Red와 Blue가 동시에 빠지지 않았으면 정답
			if board[rx][ry] == 'O' && board[rx][ry] != board[bx][by] {
				return ret
			}

			for i := 0; i < 4; i++ {
				crx, cry, cbx, cby := rx, ry, bx, by

				// Red 이동
				for board[crx+dx[i]][cry+dy[i]] != '#' && board[crx][cry] != 'O' {
					crx += dx[i]
					cry += dy[i]
				}

				// Blue 이동
				for board[cbx+dx[i]][cby+dy[i]] != '#' && board[cbx][cby] != 'O' {
					cbx += dx[i]
					cby += dy[i]
				}

				// 두 구슬의 이동 후 위치가 같다면, 기존 위치 - 이동 위치의 절대값을 통해 위치 조정
				if crx == cbx && cry == cby {
					if board[crx][cry] == 'O' {
						continue
					}

					if (math.Abs(float64(crx-rx)) + math.Abs(float64(cry-ry))) < (math.Abs(float64(cbx-bx)) + math.Abs(float64(cby-by))) {
						cbx -= dx[i]
						cby -= dy[i]
					} else {
						crx -= dx[i]
						cry -= dy[i]
					}
				}

				current_s := fmt.Sprintf("%d%d%d%d", crx, cry, cbx, cby)
				if visited[current_s] == true {
					continue
				}

				// fmt.Println("current queue", crx, cry, cbx, cby)

				q.Append(Item{Red: location{crx, cry}, Blue: location{cbx, cby}})
				visited[current_s] = true
			}
		}
		if ret == 10 {
			return -1
		}
		ret++
	}
	return -1
}

func main() {
	var N, M int

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
	}

	res := bfs(board)
	fmt.Println(res)
}
