/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// @lc code=end

/*
如果将数组 nums 中的所有元素按照单调递增或单调递减的顺序排序，那么中间下标一定是众数。
*/