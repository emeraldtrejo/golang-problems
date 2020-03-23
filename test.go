package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("Enter a string : ")
	fmt.Scanf("%s", &s)

	if (strings.Index(s, "i") > 0) && (strings.Index(s, "a") > 0) && (strings.Index(s, "n") > 0) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}
}
