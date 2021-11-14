/**
 * Write a function called someRecursive which accepts an array and a callback.
 * The function returns true if a single value in the returns true when passed
 * to the callback. Otherwise returns false.
 */
package main

import "fmt"

type fn func(int) bool

func someRecursive(array []int, callback fn) bool {
	if len(array) == 0 {
		return false
	}
	return callback(array[0]) || someRecursive(array[1:], callback)
}

func isOdd(val int) bool {
	return val%2 != 0
}

func main() {
	fmt.Println(someRecursive([]int{1, 2, 3, 4}, isOdd))
	fmt.Println(someRecursive([]int{2, 4}, isOdd))

}
