package main

import (
	"fmt"
)

func main() {
	var PI float32 = 3.14
	fmt.Println(PI)
	var n int = 0
	for n < 10 {
		n++
		fmt.Println(n)
		if n == 5 {
			break //continue
		}
	}
}
