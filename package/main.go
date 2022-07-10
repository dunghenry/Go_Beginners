package main

import (
	"fmt"
	"video/controller/demo"
	"video/controller/test"
)

func main() {
	fmt.Println("Starting")
	demo.Hi()
	fmt.Println(test.New("Hehe"))
}