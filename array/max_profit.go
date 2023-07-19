package array

import "fmt"

/*
	题目: 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
		设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
		注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // dp[i][0] 表示第i天不持有股票所得最多现金
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1]) // dp[i][i] 持有股票 所得的最多现金
	}

	return dp[len(prices)-1][0]

}

func MaxProfitUp(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}
	dp[0][0] = 0          // 没有操作 （其实我们也可以不设置这个状态）
	dp[0][1] = -prices[0] // 第一次持有股票
	dp[0][2] = 0          // 第一次不持有股票
	dp[0][3] = -prices[0] // 第二次持有股票
	dp[0][4] = 0          // 第二次不持有股票
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	fmt.Println(dp)
	return dp[len(prices)-1][4]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
