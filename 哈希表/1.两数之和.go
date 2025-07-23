/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	//哈希表
	var res map[int]int = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := res[nums[i]]; ok {
			return []int{res[nums[i]], i}
		} else {
			res[target-nums[i]] = i
		}
	}
	return nil
}

// @lc code=end

