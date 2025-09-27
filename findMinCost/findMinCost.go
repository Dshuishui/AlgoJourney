package main

import "fmt"

// Find the minimum cost that the letter slice array can be transformde to  non-decreasing order

// https://ac.nowcoder.com/acm/contest/11211/A

func minCost(n int, s string) int {
	dp := make([][]int, n)
	dp[0] = make([]int, 26)

	// initialize the first row, from A to Z， base case
	for j := range 26 {
		targetChar := byte('A' + j)
		if s[0] == targetChar {
			dp[0][j] = 0
		} else {
			dp[0][j] = 1
		}
	}

	for i := 1; i < n; i++ {
		dp[i] = make([]int, 26)
		// initialize the cost to a large number
		lastLineCost := n + 1
		for j := range 26 {
			// find the minimum cost in the last line
			if dp[i-1][j] < lastLineCost {
				lastLineCost = dp[i-1][j]
			}

			targetChar := byte('A' + j)
			currentCost := 0
			if s[i] != targetChar {
				currentCost = 1
			}
			dp[i][j] = lastLineCost + currentCost
		}
	}
	finalCost := n + 1
	for j := range 26 {
		if dp[n-1][j] < finalCost {
			finalCost = dp[n-1][j]
		}
	}
	return finalCost
}

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	fmt.Println(minCost(n, s))
}
