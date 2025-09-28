package main

import "fmt"

func lenLongestFibSubseq(arr []int) int {
	return 0
}

func main() {
	var n int
	fmt.Scan(&n)
	var arr []int
	var input int
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		arr = append(arr, input)
	}
	fmt.Println(lenLongestFibSubseq(arr))
}
