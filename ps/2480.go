package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var check [6]int

	fmt.Scanf("%d %d %d\n", &a, &b, &c)

	check[a-1]++
	check[b-1]++
	check[c-1]++

	res := 0
	for i := 0; i < len(check); i++ {
		if check[i] == 3 {
			fmt.Println(10000 + (1000 * (i + 1)))
			return
		} else if check[i] == 2 {
			fmt.Println(1000 + (100 * (i + 1)))
			return
		} else if check[i] == 1 {
			res = 100 * (i + 1)
		}
	}

	fmt.Println(res)
}
