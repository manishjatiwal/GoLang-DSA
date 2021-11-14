/**
 * Write a recursive function called isPalindrome which returns true is the
 * given strig is palindrome.
 */
package main

import "fmt"

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] == s[len(s)-1] {
		return isPalindrome(s[1 : len(s)-1])
	}
	return false
}

func main() {
	fmt.Println(isPalindrome("manish"))
	fmt.Println(isPalindrome("jyotiitoyj"))
}
