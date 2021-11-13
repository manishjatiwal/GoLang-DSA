/**
 * Write a function called averagePair. Given a sorted array of
 * integers and a target average, determine if there is a pair of values in
 * the array where the average of the pair equals the target average.
 * There may be more than one pair that matches the average target.
 */
package main

import "fmt"

func averagePair(array []int, avg int) [2]int {
	length := len(array)
	var i int = 0
	var j int = length - 1
	for i < j {
		average := (array[i] + array[j]) / 2
		if average == avg {
			result := [2]int{array[i], array[j]}
			return result
		}
		if average < avg {
			i += 1
		} else {
			j += 1
		}
	}
	return [2]int{-1, -1}
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array, 3)
}
