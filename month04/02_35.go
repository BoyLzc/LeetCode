package month04

import "fmt"

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。
*/
// 情况 1：插入元素在所有元素之前
// 情况 2：插入元素在所有元素之中
// 情况 3：插入元素在所有元素之后
// 情况 4：插入元素数组中
// 暴力解法

func searchInsert(nums []int, target int) int {
	// 处理情况1
	if target < nums[0] {
		return 0
	} else if target > nums[len(nums)-1] { // 处理情况3
		return len(nums)
	} else {
		for i := 0; i < len(nums); i++ {
			if target == nums[i] {
				return i
			} else if target > nums[i] && target < nums[i+1] {
				return i + 1
			}
		}
	}
	return -1
}

func searchInsert2(nums []int, target int) int {
	// 1.2.4种情况
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	// 3
	return len(nums)
}

// 利用二分法
func searchInsert3(nums []int, target int) int {
	n := len(nums)
	// 左闭右闭
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 情况 1、2、3
	return right + 1
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}
