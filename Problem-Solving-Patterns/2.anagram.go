/**
 * Given two strings, write a function that determines if the second string is
 * anagram of the first. An anagram is a word, phrase, or name formed by
 * rearranging the  letter of another, such as cineme, formed form iceman.
 */
package main

import (
	"fmt"
)

func validAnagram(string1 string, string2 string) bool {
	var length1 int = len(string1)
	var length2 int = len(string2)
	if length1 != length2 {
		return false
	}
	frequencyCounter := make(map[byte]int)
	for i := 0; i < length1; i++ {
		char := string1[i]
		frequencyCounter[char] = frequencyCounter[char] + 1
	}
	for i := 0; i < length2; i++ {
		char := string2[i]
		frequencyCounter[char] = frequencyCounter[char] - 1
	}
	for _, value := range frequencyCounter {
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	var isValidAnagram bool = validAnagram("ann", "nna")
	fmt.Println(isValidAnagram)
}
