/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	ans := ""
	if len(strs) == 0 {
		return ans	
	}
	index := 0
	for index < len(strs[0]) {
		tmp := strs[0][index]
		for _, x := range strs {
			if index >= len(x) || tmp != x[index] {
				return ans
			}
		}
		ans += string(tmp)
		index++
	}
	return ans
}
// @lc code=end

