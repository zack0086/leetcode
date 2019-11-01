/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	flag := false
    for i, x := range nums {
		if !flag && x == target {
			ans[0] = i
			flag = true
		}
		if flag && x != target {
			ans[1] = i - 1
			flag = false
		}
	}
	if ans[0] != -1 && ans[1] == -1 {
		ans[1] = len(nums) - 1
	}
	return ans
}
// @lc code=end

