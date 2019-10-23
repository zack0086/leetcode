/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	const MaxInt = int(^uint(0) >> 1)
	l := len(nums)
	sum := 0
    if l <= 3 {
		for _, x := range nums {
			sum += x
		}
		return sum
	}
	sort.Ints(nums)
	mindiff := MaxInt
	minsum := nums[0] + nums[1] + nums[2]
	for i := 0; i < l - 2; i++ {
		for j := i + 1; j < l - 1; j++ {
			for k := j + 1; k < l; k++ {
				sum := nums[i] + nums[j] + nums[k]
				diff := sum - target
				if abs(diff) < mindiff {
					mindiff = abs(diff)
					minsum = sum
				}
				if diff > 0 {
					break
				}
			}
		}
	}
	return minsum
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
// @lc code=end

