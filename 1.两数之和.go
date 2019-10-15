/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums{
		if _, ok := m[val]; ok {
			ans := []int{m[val], i}
			return ans
		} else {
			m[target-val] = i
		}
	}
	return nil
}
// @lc code=end

