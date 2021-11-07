/**
 * Write a function called isSubsequence which takes in two strings and checks
 * whether the characters in the first string form a subsequence of the
 * characters in the second string. In other words, the function should check
 * whether the characters in the first string appear somewhere in the second
 * string, without their order changing.
 */
package main

import "fmt"

func isSubsequence(s1 string, s2 string) bool {
	var i int = 0
	var j int = 0
	if s1 == "" {
		return true
	}
	for j < len(s2) {
		if s1[i] == s2[j] {
			i += 1
		}
		if i == len(s1) {
			return true
		}
		j += 1
	}
	return false
}

func main() {
	fmt.Println(isSubsequence("hllo", "hxello Mains"))
}
