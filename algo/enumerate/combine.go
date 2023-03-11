package enumerate

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合
// 可以使用递归做，下面是一种比较快的做法
func combine(n int, k int) (ans [][]int) {
	temp := []int{}
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1)

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		temp[j]++
	}
	return
}
