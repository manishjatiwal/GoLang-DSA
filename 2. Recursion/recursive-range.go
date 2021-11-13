/**
 * Write a function called recursiveRange which accepts a number and adds up
 * all the numbers from 0 to that number
 */
package main

import "fmt"

func recursiveRange(number int) int {
	if number <= 0 {
		return 0
	}
	return number + recursiveRange(number-1)
}

func main() {
	fmt.Println(recursiveRange(6))
	fmt.Println(recursiveRange(10))
	fmt.Println(recursiveRange(-10))
}
