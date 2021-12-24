package main

import (
	"fmt"
)

func getGCD(A, B int) int {
	if B == 0 {
		return A
	} else {
		return getGCD(B, A%B)
	}
}

func getLCM(A, B int) int {
	gcd := getGCD(A, B)
	return (A * B) / gcd
}

func main() {
	var T int
	fmt.Scan(&T)

	for ; T > 0; T-- {
		var A, B int
		fmt.Scanf("%d %d\n", &A, &B)

		fmt.Println(getLCM(A, B))
	}
}
