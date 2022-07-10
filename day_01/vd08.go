package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}
	for index, item := range nums {
		fmt.Println(index, item)
	}

	for index, item := range nums {
		if item == 3 {
			fmt.Println(index, item)
		}
	}

	mapstr := map[string]string{"1": "Mango", "2": "Apple"}
	for k, v := range mapstr {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range mapstr {
		fmt.Println("key:", k)
	}

	for i, c := range "Hi" {
		fmt.Println(i, c)
	}
}
