// Binary Search
package main

import "fmt"

func binarySearch(array []int, n int) int {
	var min int = 0
	var max int = len(array) - 1
	var index int = -1
	for min < max {
		index = (min + max) / 2
		if n == array[index] {
			return index
		}
		if n < array[index] {
			max = index - 1
		} else {
			min = index + 1
		}
	}
	return index
}

func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(array, 3))
}
