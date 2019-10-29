/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	var ans [][]int
	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	i, j := 0, 1
	for {
		tgt := -nums[i] - nums[j]
		for k := j + 1; k < len(nums); k++ {
			if nums[k] == tgt {
				tmp := []int{nums[i], nums[j], nums[k]}
				ans = append(ans, tmp)
				break
			}
		}
		curi, curj := nums[i], nums[j]
		if j == len(nums) - 1 {
			for nums[i] == curi && i < len(nums) - 2 {
				i++
			}
			j = i + 1
			if nums[i] + nums[j] > 0 {
				break
			}
		} else {
			for nums[j] == curj && j < len(nums) - 1 {
				j++
			}
		}
		if i >= len(nums) - 2 {
			break
		}
	}
	return ans
}
// @lc code=end

