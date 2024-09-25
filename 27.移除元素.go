/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	var arr []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			arr = append(arr, nums[i])
		}
	}
	copy(nums, arr)
	return len(arr)
}

// @lc code=end

