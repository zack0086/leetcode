/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 1
	for i < len(nums) {
		if nums[i] == nums[i-1] {
			nums = append(nums[0:i], nums[i+1:len(nums)]...)
		} else {
			i++
		}
	}
	return len(nums)
}
// @lc code=end

