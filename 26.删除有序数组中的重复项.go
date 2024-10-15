/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	//快慢指针
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] { //判断快指针是否在向下移动时，慢指针符合条件
			nums[slow] = nums[fast]
			slow++ //慢指针移动
		}
	}
	return slow

}

// @lc code=end

