package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var T int
	var res float32
	in := bufio.NewScanner(os.Stdin)

	fmt.Scan(&T)

	for ; T > 0; T-- {
		res = 0
		in.Scan()

		vec := strings.Split(in.Text(), " ")
		for _, idx := range vec {
			switch idx {
			case "@":
				res *= 3.00
			case "%":
				res += 5.00
			case "#":
				res -= 7.00
			default:
				v, _ := strconv.ParseFloat(idx, 32)
				res += float32(v)
			}
		}

		fmt.Printf("%.2f\n", res)
	}
}
