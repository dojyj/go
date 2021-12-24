package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b, op, res string

	fmt.Scanf("%s\n%s\n%s", &a, &op, &b)

	alen := len(a)
	blen := len(b)

	switch op {
	case "+":
		if alen > blen {
			res = a[0 : alen-blen]
			res += b
		} else if alen < blen {
			res = b[0 : blen-alen]
			res += a
		} else {
			res = strings.Replace(a, "1", "2", -1)
		}
	case "*":
		res = "1"
		for i := 0; i < (alen-1)+(blen-1); i++ {
			res += "0"
		}
	}

	fmt.Println(res)
}
