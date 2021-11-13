/**
* Write a function called sameFrequency. Given two positive integers,
* find out if the two numbers have the same frequency of digits.
 */
package main

import "fmt"

func sameFrequency(i int, j int) bool {
	frequencyMap := make(map[int]int)
	for i > 0 {
		var quotient int = i % 10
		i /= 10
		frequencyMap[quotient] += 1
	}
	for j > 0 {
		var quotient int = j % 10
		j /= 10
		frequencyMap[quotient] -= 1
		if frequencyMap[quotient] < 0 {
			return false
		}
	}
	fmt.Println(frequencyMap)
	return true
}

func main() {
	fmt.Println(sameFrequency(1012, 2011))
}
