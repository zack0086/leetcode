/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	maxlen := 0
	curlen := 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = i
			curlen++
		} else {
			i = m[s[i]]
			if curlen > maxlen {
				maxlen = curlen
			}
			curlen = 0
			m = make(map[byte]int)
		}
	}
	if curlen > maxlen {
		maxlen = curlen
	}
	return maxlen
}
// @lc code=end

