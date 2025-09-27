package main

import (
	"fmt"
	"math/rand"
)

// Find Kth largest element in an array use quick select

// https://leetcode.cn/problems/kth-largest-element-in-an-array/description/?envType=problem-list-v2&envId=sorting

func findKthLargest(nums []int, k int) int {
	// boundary case
	if len(nums) == 0 || k > len(nums) {
		return -1
	}
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left int, right int, k int) int {
	// when the array contains only one element, return the element
	if left == right {
		return nums[left]
	}

	middle := partition(nums, left, right)
	if middle == k {
		return nums[middle]
	} else if middle > k {
		return quickSelect(nums, left, middle-1, k)
	} else {
		return quickSelect(nums, middle+1, right, k)
	}
}

// this partition function is the same as the one in quick sort which use the single pointer method
func partition(nums []int, left int, right int) int {
	pivotRandom := rand.Intn(right-left+1) + left
	swap(nums, pivotRandom, right)
	pivot := nums[right]

	slow, fast := left, left
	for fast < right {
		if nums[fast] < pivot {
			swap(nums, slow, fast)
			slow++
		}
		fast++
	}
	swap(nums, slow, right)
	return slow
}

// this partition function use the double pointer method which is more efficient than the single pointer method
func partition_double_pointer(nums []int, left int, right int) int {
	pivotRandomIndex := left + rand.Intn(right-left+1)
	swap(nums, pivotRandomIndex, left)
	pivot := nums[left]

	less, greater := left+1, right
	for less <= greater {
		for less <= greater && nums[less] < pivot {
			less++
		}
		for less <= greater && nums[greater] > pivot {
			greater--
		}
		if less >= greater {
			break
		}
		swap(nums, less, greater)
		less++
		greater--
	}
	swap(nums, greater, left)
	return greater
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
