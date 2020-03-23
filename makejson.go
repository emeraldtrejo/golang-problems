package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Addr string
}

func main() {

	var n string
	var a string
	fmt.Printf("Enter name:")
	//user input
	fmt.Scan(&n)
	fmt.Printf("Enter address:")
	fmt.Scan(&a)

	p1 := map[string]string{
		"Name":    n,
		"Address": a,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	// Convert bytes to string.

	fmt.Println(string(b))
}
