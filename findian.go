package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter text:")
	//user input
	s, _ = reader.ReadString('\n')
	word := strings.Join(strings.Fields(s), "")
	words := strings.ToLower(word)
	if strings.HasPrefix(words, "i") && strings.HasSuffix(words, "n") && strings.Contains(words, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
