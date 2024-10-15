/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	min := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			max = maxNumber(max, prices[i]-min)
		}
	}
	return max
}

func maxNumber(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end