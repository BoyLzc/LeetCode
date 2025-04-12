package month04

/*
给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。

不能使用任何内置的库函数，如  sqrt 。
*/

func isPerfectSquare(num int) bool {
	var result bool = false
	var left, right int = 1, num
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == num {
			result = true
			return result
		} else if mid <= num/mid { // mid 小于 平方根
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func main() {

}
