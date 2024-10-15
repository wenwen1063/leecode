/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	//双指针，快慢指针
	if len(nums) == 0 {
		return 0
	}
	slow := 1 //慢指针从一开始，第一个元素固定不判断
	num := 1  //记录重复次数
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
			num = 1
		} else if nums[fast] == nums[fast-1] && num == 1 {
			nums[slow] = nums[fast]
			slow++
			num = 0
		}
	}
	return slow
}

// @lc code=end
