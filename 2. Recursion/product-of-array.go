/**
 * Write a function called productOfArray which takes an array of integers and
 * returns the product of them all.
 */

package main

import "fmt"

func productOfArray(array []int) int {
	length := len(array)
	if length == 0 {
		return 1
	}
	last := array[length-1]
	slice := array[:length-1]
	return last * productOfArray(slice)
}

func main() {
	fmt.Println(productOfArray([]int{1, 2, 3}))
	fmt.Println(productOfArray([]int{1, 2, 3, 10}))
}
