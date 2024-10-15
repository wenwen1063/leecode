/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	max := nums[0] //可以到达的最大坐标
	for k, v := range nums {
		if k+v > max {
			max = k + v
		} else {
			if max <= k && max != (len(nums)-1) { //判断是否能否到达最后一个坐标
				return false
			}
		}

	}
	return true
}

// @lc code=end

