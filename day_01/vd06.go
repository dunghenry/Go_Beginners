package main

import "fmt"

func main() {
	//Note fallthrough
	var day int
	fmt.Println("Nhap thu so tu 1 den 7")
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Println("Chu nhat")
		break
	case 2:
		fmt.Println("Thu hai")
		break
	default:
		fmt.Println("Khong hop le")
	}
}
