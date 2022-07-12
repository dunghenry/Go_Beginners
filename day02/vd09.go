package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("length: %d, capacity: %d, %v \n", len(s), cap(s), s)
}

func main() {
	slice := make([]int, 10)
	printSlice(slice)
	slice1 := make([]int, 0, 10)
	printSlice(slice1)
	slice2 := slice1[:5]
	printSlice(slice2)
	slice3 := slice2[2:5]
	printSlice(slice3)
}
