package month04

/*
	x 的平方根 给定非负数 计算 x的算术平方根
	不允许使用内置指数函数和算符 例如 pow(x, 0.5) 或者 x ** 0.5
*/

// 暴力破解
func mySqrt(x int) int {
	var result int
	for i := 0; i <= x; i++ {
		if i*i <= x && (i+1)*(i+1) > x {
			result = i
			break
		}
	}
	return result
}

func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	// 设定边界 [1, x]
	left, right := 1, x
	// 二分法
	for left <= right {
		mid := left + (right-left)/2
		// 如果 mid 小于 x的算数平方根
		if mid <= x/mid {
			// 往后顺延 如果 mid + 1 的平方根大于 x 则证明 mid 为 x 的算数平方根整数部分
			if (mid + 1) > x/(mid+1) {
				return mid
			} else { // 如果 mid + 1 的平方根还是 小于 x 的算数平方根 则左边界往右移动
				left = mid + 1
			}
		} else if mid > x/mid { // 如果 mid 大于 x 的算数平方根
			right = mid - 1
		}
	}
	return -1
}

func main() {

}
