/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	var ans []string
	if len(digits) == 0 {
		return ans
	}
	numMap := map[rune][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"}}
    for _, x := range digits {
		tmp := numMap[x]
		ans = combine(ans, tmp)
	}
	return ans
}

func combine(s1 []string, s2 []string) []string {
	if len(s1) == 0 {
		return s2
	}
	var ans []string
	for _, x := range s1 {
		for _, y := range s2 {
			ans = append(ans, string(x) + string(y))
		}
	}
	return ans
}
// @lc code=end

