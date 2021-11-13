/**
 * Write a recursive function called fib which accepts a number and returns the
 * nth number in the Fibonacci sequence.
 */
package main

import "fmt"

func fib(number int) int {
	if number <= 2 {
		return 1
	}
	return fib(number-1) + fib(number-2)
}

func main() {
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(10))
}
