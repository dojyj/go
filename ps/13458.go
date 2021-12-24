package main

import (
	"bufio"
	"fmt"
	"os"
)

// Fscan is much faster than Scan
func main() {
	var N, B, C, res int
	r := bufio.NewReader(os.Stdin)

	fmt.Fscan(r, &N)

	place := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &place[i])
	}
	fmt.Fscan(r, &B, &C)

	for _, v := range place {
		v -= B
		res++

		if v <= 0 {
			continue
		}
		res += v / C
		if v%C != 0 {
			res++
		}
	}

	fmt.Println(res)
}
