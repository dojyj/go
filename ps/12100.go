package main

import (
	"fmt"
)

type Queue struct {
	block chan int
}

func (q *Queue) Size() int {
	return len(q.block)
}

func (q *Queue) Push(block int) {
	q.block <- block
}

func (q *Queue) Pop() int {
	return <-q.block
}

var ans int

func tilt(board [][]int, direction int) {
	blen := len(board)
	// 1. per rows or columns, push all values to queue. if value == 0, ignore
	// 2. by indexing, update board. if cur pop value == post pop value, combine. this action occurs only once
	q := Queue{block: make(chan int, 100)}
	switch direction {
	case 0: // left
		for i := 0; i < blen; i++ {
			for j := 0; j < blen; j++ {
				q.Push(board[i][j])
				board[i][j] = 0
			}

			idx := 0
			for q.Size() > 0 {
				value := q.Pop()

				if value == 0 {
					continue
				}

				if board[i][idx] == 0 {
					board[i][idx] = value
				} else {
					if board[i][idx] == value {
						board[i][idx] = value * 2
						idx++ // combine only once
					} else {
						idx++
						board[i][idx] = value
					}
				}
			}
		}
	case 1: // right
		for i := 0; i < blen; i++ {
			for j := blen - 1; j >= 0; j-- {
				q.Push(board[i][j])
				board[i][j] = 0
			}

			idx := blen - 1
			for q.Size() > 0 {
				value := q.Pop()

				if value == 0 {
					continue
				}

				if board[i][idx] == 0 {
					board[i][idx] = value
				} else {
					if board[i][idx] == value {
						board[i][idx] = value * 2
						idx--
					} else {
						idx--
						board[i][idx] = value
					}
				}
			}
		}
	case 2: // up
		for i := 0; i < blen; i++ {
			for j := 0; j < blen; j++ {
				q.Push(board[j][i])
				board[j][i] = 0
			}

			idx := 0
			for q.Size() > 0 {
				value := q.Pop()

				if value == 0 {
					continue
				}

				if board[idx][i] == 0 {
					board[idx][i] = value
				} else {
					if board[idx][i] == value {
						board[idx][i] = value * 2
						idx++
					} else {
						idx++
						board[idx][i] = value
					}
				}
			}
		}
	case 3: // down
		for i := 0; i < blen; i++ {
			for j := blen - 1; j >= 0; j-- {
				q.Push(board[j][i])
				board[j][i] = 0
			}

			idx := blen - 1
			for q.Size() > 0 {
				value := q.Pop()

				if value == 0 {
					continue
				}

				if board[idx][i] == 0 {
					board[idx][i] = value
				} else {
					if board[idx][i] == value {
						board[idx][i] = value * 2
						idx--
					} else {
						idx--
						board[idx][i] = value
					}
				}
			}
		}
	}
}

func dfs(board [][]int, cnt int) {
	blen := len(board)
	// max value update
	if cnt == 5 {
		for i := 0; i < blen; i++ {
			for j := 0; j < blen; j++ {
				if ans < board[i][j] {
					ans = board[i][j]
				}
			}
		}
		return
	}

	// tilt and dfs
	for i := 0; i < 4; i++ {
		cpy := make([][]int, len(board))
		for i := range board {
			cpy[i] = make([]int, len(board[i]))
			copy(cpy[i], board[i])
		}

		tilt(cpy, i)
		// fmt.Println(cnt, i, cpy)
		dfs(cpy, cnt+1)
	}
}

func main() {
	var N int

	fmt.Scan(&N)

	board := make([][]int, N)
	for i := 0; i < N; i++ {
		board[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Scan(&board[i][j])
		}
	}

	dfs(board, 0)
	fmt.Println(ans)
}
