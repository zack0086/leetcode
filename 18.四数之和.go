/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
    var ans [][]int
	if len(nums) < 4 {
		return ans
	}
	sort.Ints(nums)
	i, j, k := 0, 1, 2
	for {
		for l := k + 1; l < len(nums); l++ {
			if nums[i] + nums[j] + nums[k] + nums[l] == target {
				tmp := []int{nums[i], nums[j], nums[k], nums[l]}
				ans = append(ans, tmp)
				break
			}
		}
		curi, curj, curk := nums[i], nums[j], nums[k]
		if k == len(nums) - 1 && j == len(nums) - 2 {
			for nums[i] == curi && i < len(nums) - 3 {
				i++
			}
			j = i + 1
			k = i + 2
		} else if k == len(nums) - 1 {
			for nums[j] == curj && j < len(nums) - 2 {
				j++
			}
			k = j + 1
		} else {
			for nums[k] == curk && k < len(nums) - 1 {
				k++
			}
		}
		if i >= len(nums) - 3 {
			break
		}
	}
	return ans
}
// @lc code=end

