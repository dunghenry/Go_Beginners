package main

import "fmt"

func main() {
	odd := [6]int{2, 3, 4, 6, 8, 12}
	fmt.Println(odd)

	var slice []int = odd[0:1]
	fmt.Println(slice)

	names := [4]string{"John", "Jim", "Jack", "jen"}
	fmt.Println(names)

	slice1 := names[0:2]
	slice2 := names[1:3]

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice2[0] = "ZZZ"

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(names)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
	}
	fmt.Println(s)
}
