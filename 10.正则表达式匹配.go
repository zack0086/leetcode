/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	if s == "" {
		return ifpMatch(p)
	} 
	if p == "" {
		return false
	}
	if len(p) == 1 {
		if p[0] == s[0] || p[0] == '.' {
			return isMatch(s[1:], p[1:])
		} else {
			return false
		}
	}
	if (p[0] == s[0] || p[0] == '.') && p[1] != '*' {
		return isMatch(s[1:], p[1:])
	} else if (p[0] == s[0] && p[1] == '*') || (p[0] == '.' && p[1] == '*') {
		return isMatch(s[1:], p[2:]) || isMatch(s[1:], p) || isMatch(s, p[2:])
	} else if p[1] == '*' {
		return isMatch(s, p[2:])
	} else {
		return false
	}
}

func ifpMatch(p string) bool {
	if p == "" {
		return true
	}
	if len(p) % 2 == 0 {
		for i := 1; i < len(p) ; i = i + 2 {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}
	return false
}
// @lc code=end

