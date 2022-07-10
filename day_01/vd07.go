package main

import "fmt"

func main() {
	n := 10

	for i := 0; i < n; i++ {
		fmt.Println(i)
	}

	for i := 0; i < n; i++ {
		for j := n; j > 0; j-- {
			fmt.Println(i, j)
		}
	}

	for true {
		fmt.Println("This loop will run forever.")
	}

	sum := 1

	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}
}
