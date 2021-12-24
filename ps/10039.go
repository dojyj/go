package main

import (
	"fmt"
)

func main() {
	res := 0
	for i := 0; i < 5; i++ {
		var n int
		fmt.Scan(&n)
		if n < 40 {
			n = 40
		}
		res += n
	}
	fmt.Println(res / 5)
}
