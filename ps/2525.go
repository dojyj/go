package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Scanf("%d %d\n%d", &a, &b, &c)

	min := b + c
	for min >= 60 {
		min -= 60
		a++
		if a >= 24 {
			a = 0
		}
	}

	fmt.Println(a, min)
}
