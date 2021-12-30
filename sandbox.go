package main

import (
	"fmt"
)

func main() {
	// if logic check
	var i = -1
	a := make([]int, 3)
	if a[i] == 0 && i >= 0 {
		fmt.Println(a[i])
	}
}
