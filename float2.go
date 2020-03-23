package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x float64
	fmt.Println("Enter a float value : ")
	fmt.Scanln(&x)
	var y int = int(x)
	s := strconv.Itoa(y)
	fmt.Println(s)

}
