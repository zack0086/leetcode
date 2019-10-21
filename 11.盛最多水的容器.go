/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	length := len(height)
	s := make([]int, length)
	s[0] = 0
	for i := 1; i < length; i++ {
		tmp := maxS(i, height)
		if s[i-1] > tmp {
			s[i] = s[i-1]
		} else {
			s[i] = tmp
		}
	}
	maxA := 0
	for _, x := range s {
		if x > maxA {
			maxA = x
		}
	}
	return maxA
}

func maxS(n int, height []int) int {
	tmpMax := 0
	for i := 0; i < n; i++ {
		h := height[n]
		if height[i] < h {
			h = height[i]
		}
		if h*(n-i) > tmpMax {
			tmpMax = h * (n-i)
		}
	}
	return tmpMax
}

// @lc code=end

