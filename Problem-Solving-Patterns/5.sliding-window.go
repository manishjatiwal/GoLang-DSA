/**
 * Write a function called maxSubarraySum, which accepts an array of integers
 * and a number called n. The function should calculate the maximum sub of n
 * consecutive elements in array.
 */
package main

import "fmt"

func maxSubarraySum(array []int, n int) int {
	if len(array) == 0 {
		return 0
	}
	var tempSum int = 0
	var i int = 0
	for i < n {
		tempSum += array[i]
		i += 1
	}
	var maxSum int = tempSum
	for i < len(array) {
		tempSum = tempSum + array[i] - array[i-n]
		if tempSum > maxSum {
			maxSum = tempSum
		}
		i += 1
	}
	return maxSum
}

func main() {
	var array = []int{1, 2, 5, 2, 8, 1, 5}
	fmt.Println(maxSubarraySum(array, 2))
	fmt.Println(maxSubarraySum(array, 4))
}
