/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 */

// @lc code=start
func findLHS(nums []int) int {
	//哈希表 先把所以的值统计到哈希表中，然后在判断有多少值【浮动一】
	res := map[int]int{}
	for _, v := range nums {
		res[v]++
	}
	ans := 0
	for k, v := range res {
		if c1 := res[k+1]; c1 > 0 && v+c1 > ans {
			ans = v + c1
		}
	}
	return ans
}

// @lc code=end

