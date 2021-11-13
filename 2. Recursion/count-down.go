/**
 * WAP that uses recursion to count down from a given number
 */
package main

import "fmt"

func countDown(num int) {
	fmt.Println((num))
	if num == 1 {
		return
	}
	countDown(num - 1)
	return
}

func main() {
	fmt.Println("Count Down")
	countDown(10)
}
