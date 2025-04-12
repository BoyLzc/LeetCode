package month04

/*
二分查找
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
*/

// 左闭右闭
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1 // target在左闭右闭的区间 [left, right]
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { // target 在左闭右闭 [mid + 1, right]
			left = mid + 1
		} else { // target 在左闭右闭的区间 [left, mid - 1]
			right = mid - 1
		}
	}
	return -1
}

// 左闭右开
func search2(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { // target 在左闭右开的区间 [mid + 1, right)
			left = mid + 1
		} else { // target 在左闭右开的区间 [left, mid)
			right = mid
		}
	}
	return -1
}

func main() {

}
