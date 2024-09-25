/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	var arr []int

	for i, j := 0, 0; i < m || j < n; {
		if i == m {
			arr = append(arr, nums2[j:]...)
			j++
		} else if j == n {
			arr = append(arr, nums1[i:]...)
			i++
		} else if nums1[i] < nums2[j] {
			arr = append(arr, nums1[i])
			i++
		} else {
			arr = append(arr, nums2[j])
			j++
		}
	}
	copy(nums1, arr)
}

// @lc code=end

