package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

// Length Of Longest Increasing Subsequence
// https://leetcode.cn/problems/longest-increasing-subsequence/description/?envType=problem-list-v2&envId=array

func lengthOfLIS(nums []int) int {
	// boundary case
	if len(nums) == 0 {
		return 0
	}

	// initialize dp array
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			// find the smaller number before nums[i] and update dp[i]
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	maxLength := 0
	// if dp is not empty, update to the first element
	if len(dp) > 0 {
		maxLength = dp[0]
	}

	for i := 0; i < len(dp); i++ {
		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}

	return maxLength
}

func main() {
	// input two lines, first line is n, second line is n numbers
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(lengthOfLIS(nums))


	// input a line of numbers directly
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// line := scanner.Text()
	// numStrs := strings.Split(line, " ")
	// nums = make([]int, 0, len(numStrs))
	// for _, numStr := range numStrs {
	// 	num, err := strconv.Atoi(numStr)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	nums = append(nums, num)
	// }
	// fmt.Println(lengthOfLIS(nums))
}
