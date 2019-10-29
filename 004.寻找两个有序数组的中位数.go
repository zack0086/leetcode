/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totallen := len(nums1) + len(nums2)
	sum := 0
	curnum := 0
	for i, i1, i2 := 0, 0, 0; i <= totallen / 2; i++ {
		if i1 >= len(nums1) {
			curnum = nums2[i2]
			i2++
		} else if i2 >= len(nums2) {
			curnum = nums1[i1]
			i1++
		} else if nums1[i1] <= nums2[i2] {
			curnum = nums1[i1]
			i1++
		} else if nums1[i1] > nums2[i2] {
			curnum = nums2[i2]
			i2++
		}

		if totallen % 2 == 0 && (i == totallen / 2 - 1 || i == totallen / 2) {
			sum += curnum
		} else if totallen % 2 != 0 && i == totallen / 2 {
			return float64(curnum)
		}
	}
	return float64(sum) / 2
}
// @lc code=end

