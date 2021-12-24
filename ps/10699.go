package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02"
	utc, _ := time.Parse(layout, time.Now().Format(layout))
	fmt.Println(utc.Format(layout))
}
