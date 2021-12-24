package main

import (
	"fmt"
)

func main() {
	var N, res int64
	var cnt int

	fmt.Scan(&N)

	i := 1
EXIT:
	for {
		res += int64(i)
		if res > N {
			break EXIT
		}
		i++
		cnt++
	}

	fmt.Println(cnt)
}
