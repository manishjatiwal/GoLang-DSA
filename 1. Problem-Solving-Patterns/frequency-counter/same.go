/**
* Write a function called same which accepts two arrays.
* The function should return true if every value in the first array has it's
* corrosponding value squared in the second array. The frequency of the values
* must be same.
 */
package main

import (
	"fmt"
)

func same(array1 []int, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}
	frequencyCounter1 := make(map[int]int)
	frequencyCounter2 := make(map[int]int)
	for _, item := range array1 {
		if frequencyCounter1[item] == 0 {
			frequencyCounter1[item] = 1
		} else {
			frequencyCounter1[item]++
		}
	}
	for _, item := range array2 {
		if frequencyCounter2[item] == 0 {
			frequencyCounter2[item] = 1
		} else {
			frequencyCounter2[item]++
		}
	}
	for key, value := range frequencyCounter1 {
		if frequencyCounter2[key*key] != value {
			return false
		}
	}
	return true
}

func main() {
	array1 := []int{2, 1, 3, 2}
	array2 := []int{4, 4, 1, 9}
	var isSame bool = same(array1, array2)
	fmt.Println(isSame)
}
