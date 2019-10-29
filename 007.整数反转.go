/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	const MaxInt = 2147483647
	const MinInt = -2147483648

	pos := true
	if x < 0 {
		pos = false
		x = -x
	}
	s := strconv.Itoa(x)
	ss := []rune(s)
	for i, j := 0, len(ss) - 1; i < j; i, j = i + 1, j - 1 {
		tmp := ss[i]
		ss[i] = ss[j]
		ss[j] = tmp
	}
	s = string(ss)
	n, _ := strconv.ParseInt(s, 10, 64)
	fmt.Println(MaxInt)
	if !pos {
		n = -n
		if n < int64(MinInt) {
			return 0
		}
	}
	if n > int64(MaxInt) {
		return 0
	}
	return int(n)
}
// @lc code=end

