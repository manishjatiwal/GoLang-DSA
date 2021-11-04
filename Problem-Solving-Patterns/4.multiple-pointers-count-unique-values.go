/**
 * Implement a function called countUniqueValues which accepts a sorted array,
 * and counts the unique values in the array.
 */
package main

import "fmt"

func countUniqueValues(array []int) int {
	length := len(array)
	if length < 2 {
		return length
	}
	i := 0
	j := 1
	for j < length {
		if array[i] != array[j] {
			i += 1
			array[i] = array[j]
		}
		j += 1
	}

	return i + 1
}

func main() {
	array := []int{1, 1, 1, 1, 4, 4, 7, 8, 8, 8, 8, 8}
	fmt.Println(countUniqueValues(array))
}
