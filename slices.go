package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Println("Enter integers for array")
	//fmt.Scan(&s)
	a := append(make([]int, 3))
	for s != "X" {
		fmt.Scan(&a)
		fmt.Println(&a)
	}
	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)
}
