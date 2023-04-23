package main

import "fmt"

func main() {
	var numbers = make([]int, 5, 10)
	statusOfArrays(numbers)
}

func statusOfArrays(arr []int) {
	fmt.Printf("len: %d, cap: %d, value: %v", len(arr), cap(arr), arr)
}
