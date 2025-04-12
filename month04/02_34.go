package month04

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
进阶：你可以设计并实现时间复杂度为 $O(\log n)$ 的算法解决此问题吗？

情况 1：只有一个目标值
情况 2：存在多个目标值
*/
func searchRange(nums []int, target int) []int {
	// 存储结果 目标值的开始位置以及结束位置
	var result = []int{}

	for i := 0; i < len(nums); i++ {
		// 情况 1 只存在一个目标值
		if len(nums) == 1 { // 如果只有一个元素
			if nums[i] == target {
				result = append(result, i, i)
				return result
			}
		} else { // 如果不止一个元素
			// 判断一下数组有无越界 如果 i + 1 越界
			if i+1 > len(nums)-1 {
				if nums[i] == target {
					result = append(result, i, i)
					return result
				}
			} else { // 不在最后
				// 无越界情况
				if nums[i] == target && nums[i+1] != target {
					result = append(result, i, i)
					return result
				}
			}

			// 情况 2
			if nums[i] == target && nums[i+1] == target {
				result = append(result, i)
				// 如果极端情况 第二个元素刚好是最后一个元素
				if i+1 == len(nums)-1 {
					result = append(result, i, i+1)
					return result
				}
				// 开始找到目标元素尾索引
				for j := i + 2; j < len(nums); j++ {
					// 判断一下数组越界 如果 target 在最后一个元素
					if j == len(nums)-1 {
						if nums[j] == target {
							result = append(result, j)
							return result
						} else {
							result = append(result, j-1)
							return result
						}
					}
					// 如果不在最后一个元素 则继续往后找
					if nums[j] != target {
						result = append(result, j-1)
						return result
					}
				}
			}
		}
	}
	// 如果遍历完毕没有找到目标值，则返回[-1, -1]
	result = append(result, -1, -1)
	return result
}

func searchRange2(nums []int, target int) []int {
	// 存储结果
	var result = []int{}

	for i := 0; i < len(nums); i++ {
		// 找到首个目标元素
		if nums[i] == target {
			// 如果 仅有一个目标元素
			if i == len(nums)-1 {
				result = append(result, i, i)
				return result
			}
			if i+1 == len(nums)-1 { // 如果下一个元素是最后一个元素
				if nums[i+1] == target {
					result = append(result, i, i+1)
					return result
				} else {
					result = append(result, i, i)
					return result
				}
			}
			// 如果下一个元素不是目标元素 则仅有一个目标元素
			if nums[i+1] != target {
				result = append(result, i, i)
				return result
			} else { //  i + 1 也是目标元素
				// 如果 i + 2 是最后一个元素
				if i+2 == len(nums)-1 {
					if nums[i+2] == target {
						result = append(result, i, i+2)
						return result
					} else {
						result = append(result, i, i+1)
						return result
					}
				}
				// 遍历后续元素
				// i + 2 不会越界
				for j := i + 2; j < len(nums); j++ {
					if nums[j] != target {
						result = append(result, i, j-1)
						return result
					}
					// 如果最后一个元素也是目标值
					if nums[j] == target && j == len(nums)-1 {
						result = append(result, i, j)
						return result
					}
				}
			}
		}
	}

	result = append(result, -1, -1)
	return result
}

// 二分法
func searchRange3(nums []int, target int) []int {
	rightBorder := getRightBorder(nums, target)
	leftBorder := getLeftBorder(nums, target)
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}
	// leftBorder + 1 是左边界
	// rightBorder - 1 是右边界
	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}
	// 存储结果
	return []int{-1, -1}
}

// 利用二分查找找到target右边界
// 思路 不断使左边界向右逼近
func getRightBorder(nums []int, target int) int {
	var rightBorder int = -2
	var left int = 0
	var right int = len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
			rightBorder = left
		} else {
			right = mid - 1
		}
	}
	return rightBorder
}

// 利用二分查找找到target左边界
// 思路 不断使右边界向左逼近
func getLeftBorder(nums []int, target int) int {
	var leftBorder int = -2
	var left int = 0
	var right int = len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
			leftBorder = right
		} else {
			left = mid + 1
		}
	}
	return leftBorder
}

func main() {

}
