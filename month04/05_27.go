package main

import "fmt"

/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。

假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：

更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
返回 k。
*/

// 思路：先移除 再换位
func removeElement(nums []int, val int) int {
	var count int  // 统计符合目标元素的数量
	var count2 int // 统计不符合目标元素的数量
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			fmt.Println(nums[i], "不等于", val)
			count2++
		} else {
			count++
		}
	}

	for i := 0; i < len(nums); i++ {
		// 如果是指定元素
		if nums[i] == val {
			nums = removeElement2(nums, i)
			count-- // 删去一个减少一个
			i--
		}
		if count == 0 {
			break
		}
	}
	fmt.Println(nums)
	return count2
}

func removeElement2(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1] // 用最后一个元素覆盖目标元素
	return slice[:len(slice)-1]        // 截取去掉最后一个元素
}

// 暴力破解法
func removeElement3(nums []int, val int) int {
	count := len(nums)
	for i := 0; i < count; i++ {
		if nums[i] == val {
			for j := i + 1; j < count; j++ {
				// 从后往前移
				nums[j-1] = nums[j]
			}
			i--
			count--
		}
	}
	return count
}

// 双指针法：通过一个快指针和慢指针在一个for循环下完成两个for循环工作
func removeElement4(nums []int, val int) int {
	var slowIndex int = 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}

func main() {
	// 切片
	nums := []int{3, 2, 2, 3, 2}
	//removeElement(nums, 2)
	fmt.Println(removeElement4(nums, 2))
	fmt.Println(nums)
}
