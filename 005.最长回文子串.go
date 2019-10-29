/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
    for length := len(s); length > 1; length-- {
		for i := 0; i + length <= len(s); i++ {
			if isPal(s, i, i + length - 1) {
				runes := []rune(s)
				ans := string(runes[i:i+length])
				return ans
			}
		}
	}
	return string(s[0])
}

func isPal(s string, start int, end int) bool {
	for i, j := start, end; i < j; i, j = i + 1, j - 1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
// @lc code=end

