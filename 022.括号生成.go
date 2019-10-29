/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
var ans []string

func generateParenthesis(n int) []string {
	ans = nil
	build(0, 0, n, "")
	return ans
}

func build(left int, right int, n int, s string) {
	if left == n && right == n {
		ans = append(ans, s)
	} else if left == n {
		build(left, right + 1, n, s + ")")
	} else if right != n {
		if left == right {
			build(left + 1, right, n, s + "(")
		} else if left > right {
			build(left + 1, right, n, s + "(")
			build(left, right + 1, n, s + ")")
		}
	}
}
// @lc code=end

