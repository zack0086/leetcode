/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	INT_MAX := 2147483647
	INT_MIN := -2147483648
	if dividend == 0 {
		return 0
	}
	ans := 0
	flag := 1
	if dividend > 0 && divisor < 0 {
		divisor = -divisor
		flag = -1
	} else if dividend < 0 && divisor > 0 {
		dividend = -dividend
		flag = -1
	} else if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	}
	for dividend >= divisor {
		ans++
		dividend -= divisor
	}
	if ans * flag > INT_MAX || ans * flag < INT_MIN {
		return INT_MAX
	}
	return ans * flag
}
// @lc code=end

