package main

import (
	"fmt"
)

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	// fmt.Print(n*m - 1)
	fmt.Print(n - 1 + (m-1)*n)
}
