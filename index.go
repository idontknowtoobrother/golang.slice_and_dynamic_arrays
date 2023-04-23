package main

import "fmt"

func main() {
	var numbers = make([]int, 5, 10)
	statusOfArrays(numbers)

	var dynamicNumbers []int
	dynamicNumbers = append(dynamicNumbers, 1)
	dynamicNumbers = append(dynamicNumbers, 12)
	statusOfArrays(dynamicNumbers)
}

func statusOfArrays(arr []int) {
	fmt.Printf("\nlen: %d, cap: %d, value: %v", len(arr), cap(arr), arr)
}
