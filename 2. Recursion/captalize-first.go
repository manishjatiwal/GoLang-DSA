/**
 * Write a recursive function called capitalizeFirst, given an array of strings,
 * capitalize the first letter of each string in the array.
 */
package main

import (
	"fmt"
	"strings"
)

func capitalizeFirst(array []string) []string {
	if len(array) == 0 {
		return array
	}
	first := array[0]
	first = strings.ToUpper(string(first[0])) + first[1:]

	return append([]string{first}, capitalizeFirst(array[1:])...)
}

func main() {
	fmt.Println(capitalizeFirst([]string{"manish", "kumar"}))
}
