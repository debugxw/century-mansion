package dp

import "algo/utils"

/**
 * 本质上是小区间推导大区间
 *
 * dp版本代码，最外层的循环，i为什么是n-1 -> 0，而不能反过来？
 * (i,j) 0 1  2   3   4   ...   n-2   n-1   n   n+1
 * 0     0 1  2   3   4   ...                   n+1
 * 1       1  2   3   4   ...                   n+1
 * 2          2   3   4   ...                   n+1
 * 3              3   4   ...                   n+1
 * 4                  4                         n+1
 * .                      .                     .
 * .                         .                  .
 * n-2                          n-2   n-1   n   n+1
 * n-1                                n-1   n   n+1
 * n+1
 *
 * 须从下往上算，即先算dp[n-1][n+1]：
 * 根据递推关系，算dp[i][j]时依赖的dp[i][k]和dp[k][j]，其中i<k<j。
 * 1、如果从上往下计算，依赖的dp[k][j]根本就还未算出（k比i大），比如算dp[0][3]时，依赖的dp[1][3]还是个未知数。
 * 2、从下往上就不一样，算dp[i][j]时，依赖的dp[i][k]，位于同一行左侧，已计算过；
 *                                    依赖的dp[k][j]，因为k>i，位于更下面的行，也已计算过。
 */
func maxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = utils.MaxInt(dp[i][j], sum)
			}
		}
	}
	return dp[0][n+1]
}
