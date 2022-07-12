package main

import "fmt"
//Đệ quy
func factory(num int) int {
	if num == 0 {
		return 1
	}
	return num * factory(num-1)
}
func main() {
	fmt.Println(factory(5))
	fmt.Println(24 * 5)
}
