package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	w_map := make(map[string]int)
	vec := strings.Split(s, " ")

	for _, v := range vec {
		w_map[v]++
	}

	return w_map
}

func main() {
	wc.Test(WordCount)
}
