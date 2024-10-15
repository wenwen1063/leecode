/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	k %= len(nums)
	// 翻转数组,让数组的结构符合k个元素倒叙的规则，在按照翻转k分割的俩个数组。让整个数组轮转
	reverse(nums)     // 翻转整个数组,让数组按照倒序排
	reverse(nums[:k]) // 翻转前k个
	reverse(nums[k:]) // 翻转后n-k个
}

// 翻转数组
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// @lc code=end

