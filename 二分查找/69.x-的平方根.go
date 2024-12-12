/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	//二分查找
	if x < 1 {
		return 0
	}
	l := 1
	h := x
	for l <= h {
		mid := l + (h-l)/2
		sqrt := x / mid
		if sqrt == mid {
			return mid
		} else if sqrt < mid {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return h
}

// @lc code=end

