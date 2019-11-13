/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 报数
 */

// @lc code=start
func countAndSay(n int) string {
	var ans []string
	ans = append(ans, "1")
	for i := 1; i < 30; i++ {
		ans = append(ans, parse(ans[i-1]))
	}
	return ans[n-1]
}

func parse(s string) string {
	ans := ""
	last := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == last {
			count++
		} else {
			ans = ans + strconv.Itoa(count) + string(last)
			last = s[i]
			count = 1
		}
	}
	ans = ans + strconv.Itoa(count) + string(last)
	return ans
}
// @lc code=end

