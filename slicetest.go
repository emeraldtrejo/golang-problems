package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var holder []int
	holder = make([]int, 0, 3)

	var userInput string
	fmt.Println("Please enter some numbers")
	userInput = readInput()

	for userInput != "X" {
		buffer, err := parseInput(userInput)
		if err != nil {
			fmt.Printf("%s is not a valid input\n", userInput)
		} else {
			holder = append(holder, buffer...)
			sort.Ints(holder)
			currentState := toString(holder)
			fmt.Printf("Entered numbers: %s\n", currentState)
		}
		userInput = readInput()
	}

}

func readInput() string {

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return "X"

}

func parseInput(userInput string) ([]int, error) {
	buffer := make([]int, 0, 3)
	splittedInput := strings.Split(userInput, " ")
	for i := range splittedInput {
		var item string
		item = splittedInput[i]
		intInput, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		} else {
			buffer = append(buffer, intInput)
		}
	}
	return buffer, nil

}

func toString(input []int) string {
	output := []string{}
	for i := range input {
		v := input[i]
		text := strconv.Itoa(v)
		output = append(output, text)
	}
	return strings.Join(output, ", ")
}
