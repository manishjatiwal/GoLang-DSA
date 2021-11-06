/**
 * Write a function called sumZero, which accepts a sorted array of intgers. The
 * function shuould find the first pair where the sum is 0. Return an array that
 * includes both the values that sum to zero or undefined if a pair doesnot
 * exist.
 */

package main

import "fmt"

func sumZero(array []int) [2]int {
	var res [2]int
	length := len(array)
	i := 0
	j := length - 1
	for i < j {
		sum := array[i] + array[j]
		if sum < 0 {
			i += 1
		} else if sum == 0 {
			res[0] = array[i]
			res[1] = array[j]
			return res
		} else {
			j -= 1
		}
	}
	return res
}

func main() {
	array := []int{-4, -3, -2, -1, 0, 1, 2, 3, 7, 10}
	fmt.Println(sumZero(array))
}
