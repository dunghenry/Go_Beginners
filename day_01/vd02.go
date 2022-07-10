package main

import "fmt"

func main() {
	//Data Types
	var i int
	// i := 1
	var f float64
	//f := 1.4
	var b bool
	//b := true
	var s string
	//s := "hello world"
	fmt.Printf("%T %T %T %T \n", i, f, b, s) // int float64 bool string
	fmt.Printf("%v %v %v %v", i, f, b, s) // 0 0 false ""
}