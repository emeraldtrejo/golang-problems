//Write a Bubble Sort program in Go. The program should prompt the user to
//type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest.
// Use your favorite search tool to find a description of how the bubble sort algorithm works.

// As part of this program, you should write a function
// called BubbleSort() which takes a slice of integers as an argument and returns nothing.
// The BubbleSort() function should modify the slice so that the elements are in sorted order.

// A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
//You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers
//and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of
//the slice in position i with the contents in position i+1.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello")
	var slice []int
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Println("Please enter  an integer")
		text, _ := reader.ReadString('\n')
		result := strings.Split(strings.TrimSpace(text), "\n")
		val, _ := strconv.Atoi(result[0])
		slice = append(slice, int(val))
	}
	fmt.Println("your raw array", slice)
	// myslice := []int{90, 89, 65, 18, 9, 23, 2131, 324, 87, 5}
	var noSwaps bool = false
	for noSwaps != true {
		noSwaps = bubbleSort(slice)
	}
	fmt.Println("Your sorted array ", slice)
	// fmt.Println(myslice)

}
func bubbleSort(sl []int) bool {
	noSwaps := true
	for i := 0; i < len(sl)-1; i++ {
		if sl[i] > sl[i+1] {

			swap(sl, i)
			noSwaps = false
		}
	}
	return noSwaps
}

func swap(sl []int, idx1 int) {
	tmp := sl[idx1]
	sl[idx1] = sl[idx1+1]
	sl[idx1+1] = tmp
}
