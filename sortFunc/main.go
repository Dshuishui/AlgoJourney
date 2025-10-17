package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 2, 8, 3, 1}
	sort.Ints(nums)   // in-place sorting (ascending)
	fmt.Println(nums) // output: [1 2 3 5 8]

	words := []string{"banana", "apple", "cherry", "pear"}
	sort.Strings(words)
	fmt.Println(words) // output: [apple banana cherry pear]
}
