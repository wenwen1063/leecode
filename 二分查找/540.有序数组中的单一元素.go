/*
 * @lc app=leetcode.cn id=540 lang=golang
 *
 * [540] 有序数组中的单一元素
 */

// @lc code=start
func singleNonDuplicate(nums []int) int {
	//二分查找 ，判断最中间的那个临界值是不是单一元素，在那边
	l := 0
	h := len(nums) - 1
	for l < h {
		mid := l + (h-l)/2
		if mid%2 == 1 {
			mid-- //保证 l/h/mid 都在偶数位，使得查找区间大小一直都是奇数
		}
		if nums[mid] == nums[mid+1] {
			l = mid + 2
		} else {
			h = mid
		}
	}
	return nums[l]
}

// @lc code=end

