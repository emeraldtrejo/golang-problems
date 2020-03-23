package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	result, err := ioutil.ReadFile("data/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}

}
