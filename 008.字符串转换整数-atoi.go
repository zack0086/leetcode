/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	INT_MAX := 2147483647
	INT_MIN := -2147483648

	pos := true
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}
	r := []byte(str)
	if r[0] == '-' {
		pos = false
		r = r[1:]
	} else if r[0] == '+' {
		r = r[1:]
	}
	if len(r) == 0 || r[0] < '0' || r[0] > '9' {
		return 0
	}
	
	ans := getInt(r)
	if pos {
		if ans > INT_MAX {
			ans = INT_MAX
		}
	} else {
		ans = -ans
		if ans < INT_MIN {
			ans = INT_MIN
		}
	}
	return ans
}

func getInt(r []byte) int {
	var ans []byte
	for {
		if len(r) == 0 {
			return 0
		}
		if r[0] == '0' {
			r = r[1:]
		} else {
			break
		}
	}
	for _, x := range r {
		if x >= '0' && x <= '9' {
			ans = append(ans, x)
		} else {
			break
		}
	}
	i, _ := strconv.Atoi(string(ans))
	return i
}
// @lc code=end

