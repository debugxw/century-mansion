package recursive

// 给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数
// 所有的递归都可以用循环实现
func lexicalOrder(n int) []int {
	ret := make([]int, n)
	num := 1
	for i := 0; i < n; i++ {
		ret[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return ret
}
