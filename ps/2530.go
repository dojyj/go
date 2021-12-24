package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int

	fmt.Scanf("%d %d %d\n%d", &a, &b, &c, &d)

	t := d + c + (b * 60) + (a * 60 * 60)
	h := (t / (60 * 60)) % 24
	m := (t / 60) % 60
	s := t % 60
	fmt.Println(h, m, s)
}
