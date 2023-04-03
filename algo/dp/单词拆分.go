package dp

// WordBreak 力扣 139、单词拆分
// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用
func WordBreak(s string, wordDict []string) bool {
	m := map[string]struct{}{}
	for _, val := range wordDict {
		m[val] = struct{}{}
	}
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		if _, exist := m[s[0:i+1]]; exist {
			dp[i] = true
			continue
		}
		for j := 0; j < i; j++ {
			if dp[j] {
				if _, exist := m[s[j+1:i+1]]; exist {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(s)-1]
}