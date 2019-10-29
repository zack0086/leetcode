/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
    if len(s) == 0 {
		return true
	}
	if len(s) % 2 != 0 {
		return false
	}
	i := 0
	for len(s) != 0 {
		if i > 0 {
			if s[i] == ')' {
				if s[i-1] == '(' {
					s = s[0:i-1] + s[i+1:len(s)]
					i -= 2
				} else {
					return false
				}
			} else if s[i] == ']' {
				if s[i-1] == '[' {
					s = s[0:i-1] + s[i+1:len(s)]
					i -= 2
				} else {
					return false
				}
			} else if s[i] == '}' {
				if s[i-1] == '{' {
					s = s[0:i-1] + s[i+1:len(s)]
					i -= 2
				} else {
					return false
				}
			}
		}
		i++
		if i >= len(s) {
			break
		}
	}
	if len(s) != 0 {
		return false
	}
	return true
}
// @lc code=end

