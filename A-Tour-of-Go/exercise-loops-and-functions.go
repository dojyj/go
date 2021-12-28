package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	save := 0.0
	cnt := 0
EXIT:
	for {
		z -= (z*z - x) / (2 * z)
		cnt++
		if z*z == x {
			break EXIT
		} else {
			if math.Abs(save-z) > 1e-8 {
				save = z
			} else {
				break EXIT
			}
		}
	}
	fmt.Println(cnt)
	return z
}

func main() {
	num := 123123123.123123123
	fmt.Println(Sqrt(num))
	fmt.Println(math.Sqrt(num))
}
