/**
* Implenment a function called, areThereDuplicates which accepts variable number
* of argumanets, and checks whether there  are any duplicates among the
* arguments passed in.
 */
package main

import (
	"fmt"
	"sort"
)

func areThereDuplicates(args ...int) bool {
	length := len(args)
	if length <= 1 {
		return false
	}
	sort.Ints(args)

	var i int = 0
	var j int = 1
	for j < length {
		if args[i] == args[j] {
			return true
		}
		i += 1
		j += 1
	}
	return false
}

func main() {
	fmt.Println(areThereDuplicates(1, 2, 2, 4, 5))
}
