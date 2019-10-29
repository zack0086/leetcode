/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	maps := map[byte]int{'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
	ans := 0
	for i := 0; i < len(s); i++ {
		ans += maps[s[i]]
		if i == 0 {
			continue
		}
		if maps[s[i]] > maps[s[i-1]] {
			ans = ans - maps[s[i-1]] * 2
		}
	}
	return ans
}
// @lc code=end

