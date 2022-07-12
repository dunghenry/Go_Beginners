package main

import "fmt"

func sum(a int, b int) int{
	return a + b
}
func addAll(args...int) (int, int){
	finalAddValue := 0
	finalSubValue := 0
	for	_,value := range args{
		finalAddValue += value
		finalSubValue -= value
	}
	return finalAddValue, finalSubValue
}
func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(addAll(1, 2, 3, 4))
}