package main

import (
	"fmt"
	"math/rand"
)
// Find Kth largest element in an array use quick sort will timeout

// https://leetcode.cn/problems/kth-largest-element-in-an-array/description/?envType=problem-list-v2&envId=sorting

func findKthLargest(nums []int, k int) int {
	// boundary case
	if len(nums) == 0 || k > len(nums) {
		return -1
	}
	quickSort(nums, 0, len(nums)-1)

	return nums[len(nums)-k]
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	middle := partition(nums, left, right)
	quickSort(nums, left, middle-1)
	quickSort(nums, middle+1, right)
}

func partition(nums []int, left int, right int) int {
	// use a random number as pivot to avoid the worst case 
	// attention: rand.Intn(n) generates a random number in [0, n), so we need to add left to shift the range to [left, right]
	pivotRandom := rand.Intn(right-left+1) + left
	swap(nums, pivotRandom, right)

	// use the rightmost element as pivot which usually spends more time than random when the array is nearly sorted
	// so we use the random pivot above
	pivot := nums[right]

	slow, fast := left, left
	for fast < right {
		// find the smaller number than pivot, swap to the front and increment slow
		if nums[fast] < pivot {
			swap(nums, slow, fast)
			slow++
		}
		fast++
	}
	swap(nums, slow, right)
	return slow
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	var k, input, num int
	var nums []int
	fmt.Scan(&num, &k)
	for i := 0; i < num; i++ {
		fmt.Scan(&input)
		nums = append(nums, input)
	}
	fmt.Println(findKthLargest(nums, k))
}
