package main

import (
	"fmt"
	"math"
)

func main() {
	var N float64

	fmt.Scan(&N)
	side := math.Sqrt(N)
	fmt.Printf("%.8f\n", side*4)
}
