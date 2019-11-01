/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
    for i := 0; i < 9; i++ {
		tmpmap := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' && tmpmap[board[i][j]] != 0 {
				return false
			} else {
				tmpmap[board[i][j]] = 1
			}
		}
	}
	for j := 0; j < 9; j++ {
		tmpmap := make(map[byte]int)
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' && tmpmap[board[i][j]] != 0 {
				return false
			} else {
				tmpmap[board[i][j]] = 1
			}
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !ifsq(board, i, j) {
				return false
			}
		}
	}
	return true
}

func ifsq(board [][]byte, a int, b int) bool {
	tmpmap := make(map[byte]int)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[a+i][b+j] != '.' && tmpmap[board[a+i][b+j]] != 0 {
				return false
			} else {
				tmpmap[board[a+i][b+j]] = 1
			}
		}
	}
	return true
}
// @lc code=end

