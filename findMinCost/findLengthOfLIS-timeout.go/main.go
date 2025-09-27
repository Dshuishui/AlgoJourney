package main

import "fmt"

// To minimize the number of changes, we need to maximize the number of letters that remain unchanged.
// This is equivalent to finding the longest non-decreasing subsequence (LNDS) in the original string.
// The length of LNDS represents the maximum number of characters we can keep unchanged.

func minChangesToNonDecreasing(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// dp[i] represents the length of the longest non-decreasing subsequence ending at the i-th character
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1 // Initialize each subsequence length to at least 1 (itself)
	}

	// O(n^2) core loop
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// If s[j] <= s[i], then s[i] can be appended to the non-decreasing subsequence ending at s[j]
			if s[j] <= s[i] {
				// Update the value of dp[i]
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	// Find the maximum value among all dp values
	maxLength := 0
	if n > 0 {
		maxLength = 1
	}
	for _, length := range dp {
		if length > maxLength {
			maxLength = length
		}
	}

	return n - maxLength
}

func main() {
	// Test
	fmt.Println(minChangesToNonDecreasing("ACEBF")) // Output: 1
	fmt.Println(minChangesToNonDecreasing("CBA"))   // Output: 2
}
