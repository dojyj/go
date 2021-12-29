package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var visited = [501][501]bool{}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func mountain(x, y int, paper [][]int) int {
	ret := 0

	// ㅏ
	if x-1 >= 0 && x+1 < len(paper) && y+1 < len(paper[0]) {
		ret = max(ret, paper[x][y]+paper[x-1][y]+paper[x+1][y]+paper[x][y+1])
	}

	// ㅓ
	if x-1 >= 0 && x+1 < len(paper) && y-1 >= 0 {
		ret = max(ret, paper[x][y]+paper[x-1][y]+paper[x+1][y]+paper[x][y-1])
	}

	// ㅗ
	if x-1 >= 0 && y-1 >= 0 && y+1 < len(paper[0]) {
		ret = max(ret, paper[x][y]+paper[x-1][y]+paper[x][y+1]+paper[x][y-1])
	}

	// ㅜ
	if x+1 < len(paper) && y-1 >= 0 && y+1 < len(paper[0]) {
		ret = max(ret, paper[x][y]+paper[x+1][y]+paper[x][y-1]+paper[x][y+1])
	}

	return ret
}

func dfs(x, y, depth int, paper [][]int) int {
	ret := 0
	if depth == 3 {
		return paper[x][y]
	}

	visited[x][y] = true
	// fmt.Println("depth : ", depth, "visit : ", x, y)
	if x-1 >= 0 && !visited[x-1][y] {
		ret = max(ret, dfs(x-1, y, depth+1, paper))
	}
	if x+1 < len(paper) && !visited[x+1][y] {
		ret = max(ret, dfs(x+1, y, depth+1, paper))
	}
	if y-1 >= 0 && !visited[x][y-1] {
		ret = max(ret, dfs(x, y-1, depth+1, paper))
	}
	if y+1 < len(paper[0]) && !visited[x][y+1] {
		ret = max(ret, dfs(x, y+1, depth+1, paper))
	}

	visited[x][y] = false
	return ret + paper[x][y]
}

func solve(paper [][]int) int {
	ret := 0

	for i := 0; i < len(paper); i++ {
		for j := 0; j < len(paper[0]); j++ {
			ret = max(ret, mountain(i, j, paper)) // check ㅗ,ㅏ,ㅓ,ㅜ
			ret = max(ret, dfs(i, j, 0, paper))   // check polyomino that depth == 3
			// fmt.Println(i, j, ret)
		}
	}

	return ret
}

func main() {
	var N, M int

	fmt.Scanf("%d %d\n", &N, &M)

	paper := make([][]int, N)
	for i := 0; i < N; i++ {
		paper[i] = make([]int, M)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			scanner.Scan()
			paper[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

	fmt.Println(solve(paper))

}
