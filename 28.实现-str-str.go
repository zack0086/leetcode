/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
		return 0
	}
	for i, x := range haystack {
		if len(haystack) - i < len(needle) {
			return -1
		}
		if byte(x) == needle[0] {
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}
// @lc code=end

