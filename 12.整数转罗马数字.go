/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	s := ""
	pos := 1
	for num != 0 {
		dig := num % 10
		switch pos {
		case 1:
			switch dig {
			case 1:
				s = "I" + s
			case 2:
				s = "II" + s
			case 3:
				s = "III" + s
			case 4:
				s = "IV" + s
			case 5:
				s = "V" + s
			case 6:
				s = "VI" + s
			case 7:
				s = "VII" + s
			case 8:
				s = "VIII" + s
			case 9:
				s = "IX" + s
			}
		case 2:
			switch dig {
			case 1:
				s = "X" + s
			case 2:
				s = "XX" + s
			case 3:
				s = "XXX" + s
			case 4:
				s = "XL" + s
			case 5:
				s = "L" + s
			case 6:
				s = "LX" + s
			case 7:
				s = "LXX" + s
			case 8:
				s = "LXXX" + s
			case 9:
				s = "XC" + s
			}
		case 3:
			switch dig {
			case 1:
				s = "C" + s
			case 2:
				s = "CC" + s
			case 3:
				s = "CCC" + s
			case 4:
				s = "CD" + s
			case 5:
				s = "D" + s
			case 6:
				s = "DC" + s
			case 7:
				s = "DCC" + s
			case 8:
				s = "DCCC" + s
			case 9:
				s = "CM" + s
			}
		case 4:
			switch dig {
			case 1:
				s = "M" + s
			case 2:
				s = "MM" + s
			case 3:
				s = "MMM" + s
			}
		}
		num /= 10
		pos++
	}
	return s
}
// @lc code=end

