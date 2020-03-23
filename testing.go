package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type names struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Welcome ", os.Args[1])
	dat, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(dat), "\n")
	fmt.Println(len(lines), lines[17])
	namArr := make([]names, 0, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		s := strings.Fields(line)
		namArr = append(namArr, names{s[0], s[1]})
	}
	for _, val := range namArr {
		fmt.Print("First Name -", val.fname, " Last Name - ", val.lname, "\n")
	}
	// for itm => slice{
	// 	fmt.Println(itm);
	// }
}
