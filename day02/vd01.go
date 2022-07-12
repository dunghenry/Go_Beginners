package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 10
	var f float64 = 6.44
	var str1 string = "101"
	var str2 string = "10.123"
	fmt.Printf("%T\n", int(f))        //int
	fmt.Printf("%T\n", float64(i))    //float64
	fmt.Printf("%T %T\n", str1, str2) //string string
	newInt, _ := strconv.ParseInt(str1, 0, 64)
	newFloat, _ := strconv.ParseFloat(str1, 64)
	fmt.Printf("%T \n", newInt)   //int64
	fmt.Printf("%T \n", newFloat) //float64
}
