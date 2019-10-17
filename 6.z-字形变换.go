/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	} else if numRows == 2 {
		s1 := ""
		s2 := ""
		for i := 0; i < len(s); i++ {
			if i % 2 == 0 {
				s1 += string(s[i])
			} else {
				s2 += string(s[i])
			}
		}
		return s1 + s2
	}
	var e [][]byte
	curCol := 0
	for i := 0; i < len(s); i++ {
		curArr := make([]byte, numRows)
		if curCol == 0 {
			for j := 0; j < numRows; j++ {
				if i == len(s) {
					break
				}
				//fmt.Println(string(s[i]))
				curArr[j] = s[i]
				i++
			}
			i--
			curCol++
		} else {
			curArr[numRows-curCol-1] = s[i]
			curCol++
			if curCol == numRows-1 {
				curCol = 0
			}
		}
		e = append(e, curArr)
	}

	//fmt.Println(e)
	ss := ""
	for j := 0; j < numRows; j++ {
		for i := 0; i < len(e); i++ {
			if e[i][j] != 0 {
				ss += string(e[i][j])
			}
		}
	}
	return ss
}

// @lc code=end

