/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
var tempboard [][]byte
var ans [][]byte

func solveSudoku(board [][]byte) {
	tempboard = board
	backtrace(0)
	copy(board, ans)
	tempboard, ans = nil, nil
}

func backtrace(count int) {
	if count == 81 {
		for i := 0; i < 9; i++ {
			var temp []byte
			for j := 0; j < 9; j++ {
				temp = append(temp, tempboard[i][j])
			}
			ans = append(ans, temp)
		}
		return
	}
	row := count / 9
	col := count % 9
	if tempboard[row][col] == byte('.') {
		for xx := byte('1'); xx <= byte('9'); xx++ {
			tempboard[row][col] = xx
			if isPlace(count) {
				backtrace(count + 1)
			}
		}
		tempboard[row][col] = byte('.')
	} else {
		backtrace(count + 1)
	}
}

func isPlace(count int) bool {
	row := count / 9
	col := count % 9
	for j := 0; j < 9; j++ {
		if tempboard[row][j] == tempboard[row][col] && j != col {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if tempboard[j][col] == tempboard[row][col] && j != row {
			return false
		}
	}
	temprow := row / 3 * 3
	tempcol := col / 3 * 3
	for j := temprow; j < temprow + 3; j++ {
		for k := tempcol; k < tempcol + 3; k++ {
			if tempboard[j][k] == tempboard[row][col] && j != row && k != col {
				return false
			}
		}
	}
	return true
}
// @lc code=end

