package main

import "fmt"

type Employee struct {
	fname string
	lname string
}

func fullname(e Employee) string {
	return e.lname + " " + e.fname
}
func (e Employee) fullname() string {
	return e.lname + " " + e.fname
}

func main() {
	e1 := Employee{"Tran", "Dung"}
	fmt.Println(fullname(e1))
	fmt.Println(e1.fullname())
}
