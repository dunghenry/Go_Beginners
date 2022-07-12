package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("length: %d, capacity: %d, %v \n", len(s), cap(s), s)
}
func main() {
	slice1 := [6]int{2, 3, 4, 6, 8, 12}
	slice2 := slice1[2:4]
	slice3 := slice1[:2]
	slice4 := slice1[2:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	printSlice(slice2)
	printSlice(slice3)
	printSlice(slice4)
}
