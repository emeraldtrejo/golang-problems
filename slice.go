package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Println("Enter array size")
	fmt.Scan(&n)
	set(n)

}

func set(n int) {
	a := make([]int, n)
	c := "X"
	b, err := strconv.Atoi(c)

	for i, v := range a {
		fmt.Println("Enter Integer")
		fmt.Scan(&a[i], v)

		if a[i] == b && err != nil {
			break
		}
	}

	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)

}
