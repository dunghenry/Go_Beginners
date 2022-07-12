package main

import (
	"fmt"
)

func main() {
	var x [5]int
	// x := [5]int{}
	for i := 0; i < 5; i++ {
		x[i] = i + 1
	}
	fmt.Println(x)

	var a = [2][2]int{{1, 2}, {2, 3}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(a[i][j])
			fmt.Print("\t")
		}
		fmt.Println()
	}
}
