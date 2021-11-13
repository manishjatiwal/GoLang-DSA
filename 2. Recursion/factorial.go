/**
 * WAP that uses recursion to calculate the factorial of a given number
 */
package main

import "fmt"

func factorial(number int) int {
	if number == 1 {
		return number
	}
	return number * factorial(number-1)
}

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
	fmt.Println(factorial(6))
}
