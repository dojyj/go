package main

import (
	"fmt"
)

func main() {
	for {
		var a, b int

		fmt.Scanf("%d %d\n", &a, &b)

		if a == 0 && b == 0 {
			return
		}

		if a > b {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
