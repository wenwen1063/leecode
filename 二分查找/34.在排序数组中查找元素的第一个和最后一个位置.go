/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	//二分查找
	left := Search(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := Search(nums, target+1) - 1
	// if nums[right] != target {
	// 	right = -1
	// }
	return []int{left, right}
}

func Search(nums []int, x int) int {
	i, j := 0, len(nums)
	for i < j {
		h := (i + j) >> 1
		if nums[h] < x {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

// @lc code=end

