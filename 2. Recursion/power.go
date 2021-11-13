/**
 * Write a function called power which accepts a base and an exponent. The
 * function should return the power of the base to the exponent. This function
 * should mimic the functionality of Math.pow() - do not worry about negative
 * bases and exponents.
 */
package main

import "fmt"

func power(base int, exponent int) int {
	if exponent == 0 {
		return 1
	}
	return base * power(base, exponent-1)
}

func main() {
	fmt.Println(power(2, 4))
	fmt.Println(power(2, 3))
	fmt.Println(power(3, 4))
}
