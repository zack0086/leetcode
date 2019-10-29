/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
    i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:len(nums)]...)
		} else {
			i++
		}
	}
	return len(nums)
}
// @lc code=end

