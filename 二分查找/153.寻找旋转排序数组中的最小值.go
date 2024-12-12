/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	// 二分查找
	h := 0
	l := len(nums) - 1
	for h < l {
		mid := (h + l) / 2
		if nums[mid] > nums[l] {
			h = mid + 1
		} else {
			l = mid
		}
	}
	return nums[h]
}

// @lc code=end

