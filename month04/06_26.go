-package month04

/*
给你一个 非严格递增排列 的数组 nums ，
请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。
*/

// 暴力破解法
func removeDuplicates(nums []int) int {
	var count int = len(nums)
	for i := 0; i < count-1; i++ {
		// 如果遇到重复元素 则原地覆盖重复元素
		if nums[i] == nums[i+1] {
			// 覆盖
			for j := i; j < count-1; j++ {
				nums[j] = nums[j+1]
			}
			i--     // 覆盖完成以后，需要再次检查相同位置的新元素
			count-- // 删除一个元素，数组长度 - 1
		}
	}
	return count
}

// 双指针法 一快一慢两个指针来做
func removeDuplicates2(nums []int) int {
	var slowIndex int = 0 // 利用慢指针存储结果
	// 利用快指针遍历
	for fastIndex := 0; fastIndex < len(nums)-1; fastIndex++ {
		if nums[fastIndex] != nums[fastIndex+1] {
			slowIndex++
			nums[slowIndex] = nums[fastIndex + 1]
		}
	}
	return slowIndex + 1
}

func main() {

}
