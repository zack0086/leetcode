/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	var ans []int
	if len(words) == 0 {
		return ans
	}
	totalLen := len(words) * len(words[0])
	for i := 0; i + totalLen <= len(s); i++ {
		if match(s[i:i+totalLen], words) {
			ans = append(ans, i)
		}
	}
	return ans
}

func match(ss string, words []string) bool {
	wMap := make(map[string]int)
	for _, x := range words {
		wMap[x]++
	}
	len1 := len(words[0])
	for i := 0; i < len(ss); i += len1 {
		if wMap[ss[i:i+len1]] > 0 {
			wMap[ss[i:i+len1]]--
		} else {
			return false
		}
	}
	return true
}
// @lc code=end

