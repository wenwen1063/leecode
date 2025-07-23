/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func partition(nums []int, l int, r int) int {
	pivot := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= pivot {
			
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

// @lc code=end

