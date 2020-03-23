package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var n string
	fmt.Printf("Enter name of file: ")
	//user input
	fmt.Scan(&n)

	file, err := os.Open(n)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		fmt.Println(eachline)

	}
}
