package bits

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）
// 可以DFS，这里利用位运算巧妙实现
func subsets(nums []int) [][]int {
	var ret [][]int
	for mark := 0; mark < 1<<len(nums); mark++ {
		var subset []int
		for i, val := range nums {
			if mark>>i&1 == 1 {
				subset = append(subset, val)
			}
		}
		ret = append(ret, subset)
	}
	return ret
}
