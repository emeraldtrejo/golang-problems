package main

import "fmt"

func main() {

	slice := make([]int, 10)

	for i, v := range slice {
		fmt.Println("Enter Integer")
		fmt.Scan(&slice[i], v)
	}

	bubbleSort(slice)
	fmt.Println(slice)

}

func bubbleSort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
