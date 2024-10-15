/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	//贪心
	/*
		对于单独交易日:设今天价格P1、明天价格p2,则今天买入、明天卖出可赚取金额p2-pi
		(负值代表亏损)。
		对于连续上涨交易日:设此上涨交易日股票价格分别为P1,P2,.....,p,则第一天买最后一天卖收益最大,即ph-p1;等价于每
		天都买卖,即pn-p1=(p2-pi)+(ps-p2)+...+(pn-1)。
		对于连续下降交易日:则不买卖收益最大,即不会亏钱。
	*/
	ans := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1] //每天的利润
		if tmp > 0 {                   //利润大于0,小于0的利润不算
			ans += tmp
		}
	}
	return ans
}

// @lc code=end

