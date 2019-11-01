/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int)  {
    if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 2; i >= 0; i--{
		if nums[i] < nums[i+1] {
			minind := i + 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j] <= nums[i] {
					break
				}
				minind = j
			}
			tmp := nums[minind]
			nums[minind] = nums[i]
			nums[i] = tmp
			sort.Ints(nums[i+1:])
			return
		}
	}
	sort.Ints(nums)
}
// @lc code=end

