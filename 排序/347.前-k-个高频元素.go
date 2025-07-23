/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	//桶排序
	//统计每个元素的出现次数
	midMap := map[int]int{}
	maxValue := 0
	for _, num := range nums { //出现次数的map
		midMap[num]++
		maxValue = max(maxValue, midMap[num])
	}
	//把出现次数相同的数放到一个桶中
	buckets := make([][]int, maxValue+1)
	for key, value := range midMap {
		buckets[value] = append(buckets[value], key)
	}
	// 倒序遍历 buckets，把出现次数前 k 大的元素加入答案
	ans := make([]int, 0, k)
	for i := maxValue; i >= 0 && len(ans) < k; i-- {
		ans = append(ans, buckets[i]...)
	}
	return ans
}

// @lc code=end

