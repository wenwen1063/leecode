/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	//贪心
	lenth := len(nums)
	end := 0
	maxPos := 0
	step := 0
	for i := 0; i < lenth-1; i++ {
		//判断每个参数最远到达的位置
		//判断之前到达的位置和当前的位置的最远距离
		//确认那个能到达最远
		maxPos = max(nums[i]+i, maxPos)
		if i == end {
			end = maxPos //当前位置更新为最远位置
			step++       //更新步数
		}
	}
	return step

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

