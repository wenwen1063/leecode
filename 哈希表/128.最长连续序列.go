/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	//哈希表
	res := map[int]bool{}
	for _, v := range nums {
		res[v] = true
	}

	ans := 0
	for k, _ := range res {
		if !res[k-1] { //判断前面是否存在，确定是否是连续中第一个
			count := 1   //计数
			current := k //判断当前是那个数
			for res[current+1] {
				count++
				current++
			}
			if count > ans {
				ans = count
			}
		}
	}
	return ans
}

// @lc code=end

