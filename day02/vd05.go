package main

import "fmt"

//Closure
func main() {
	number := 10
	squareNum := func() int {
		number *= number
		return number
	}
	fmt.Println(squareNum())
	fmt.Println(squareNum())
}
