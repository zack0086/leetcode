/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	start, maxlen := 0, 0
	var ss []int
	for i, x := range s {
		switch x {
		case '(':
			ss = append(ss, i)
		case ')':
			if len(ss) == 0 {
				start = i + 1
			} else {
				ss = ss[:len(ss)-1]
				if len(ss) == 0 {
					if i - start + 1 > maxlen {
						maxlen = i - start + 1
					}
				} else {
					if i - ss[len(ss)-1] > maxlen {
						maxlen = i - ss[len(ss)-1]
					}
				}
			}
		}
	}
	return maxlen
}
// @lc code=end

