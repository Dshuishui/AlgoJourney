package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	var stack []int
	stack = append(stack, nums[0])
	maxLength := 1
	for i := 1; i < len(nums); i++ {
		topStack := stack[len(stack)-1]
		// if the current number is greater than the top of the stack, push it onto the stack
		if nums[i] > topStack {
			stack = append(stack, nums[i])
			if len(stack) > maxLength {
				maxLength = len(stack)
			}
			// if the current number is not greater than the top of the stack, reset the stack
		} else {
			stack = stack[:0]
			stack = append(stack, nums[i])
		}
	}
	return maxLength
}

func main() {
	var num int
	fmt.Scan(&num)
	var nums []int
	var input int
	for i := 0; i < num; i++ {
		fmt.Scan(&input)
		nums = append(nums, input)
	}
	fmt.Println(findLengthOfLCIS(nums))
}
