/**
 * Write a recursive functionn called reverse which accepts a string and
 * returns a reversed string.
 */
package main

import "fmt"

func reverse(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}

	return string([]rune(s)[length-1]) + reverse(s[:length-1])
}

func main() {
	fmt.Println(reverse("Manish"))
}
