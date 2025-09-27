package main

import (
	"fmt"
)

// Find the minimum operations to reduce x to zero

// We can divide the array into three parts: left, middle, right
// left + right = x
// So we can convert the problem to find the longest subarray whose sum is sum(nums) - x

// https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/description/

func minOperations(nums []int, x int) int {
	var totalSize int
	for _, value := range nums {
		totalSize += value
	}

	target := totalSize - x

	// boundary case
	if x > totalSize {
		return -1
	} else if x == totalSize {
		return len(nums)
	}

	left := 0
	currentSum := 0
	maxLength := -1
	for right, value := range nums {

		// include the right element to sum of the current window
		currentSum += value

		for currentSum > target && left <= right {
			currentSum -= nums[left]
			left++
		}

		if currentSum == target {
			currentLength := right - left + 1
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}
	if maxLength == -1 {
		return -1
	}
	return len(nums) - maxLength
}

func main() {
	fmt.Println(minOperations([]int{3, 2, 20, 1, 1, 3}, 10))
}
