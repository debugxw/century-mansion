package dp

// 力扣 72、编辑距离   ！！！dp[i][j]的定义一定要先搞清楚！！！
// 给你两个单词word1 和word2， 请返回将word1转换成word2 所使用的最少操作次数
// 你可以对一个单词进行 增、删、改 三种操作
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 || n == 0 {
		return m + n
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}

	minInt := func(min int, arr ...int) int {
		for _, val := range arr {
			if val < min {
				min = val
			}
		}
		return min
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minInt(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}
	return dp[m][n]
}