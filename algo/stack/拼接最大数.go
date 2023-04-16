package stack

// 力扣 321、拼接最大数
// 给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字
// 现在从这两个数组中选出 k (k <= m + n) 个数字拼接成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序
// 求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	var ret []int
	for cnt1 := 0; cnt1 <= k && cnt1 <= len(nums1); cnt1++ {
		cnt2 := k - cnt1
		if cnt1 > len(nums1) || cnt2 > len(nums2) {
			continue
		}
		temp1 := removeDigits(nums1, len(nums1)-cnt1)
		temp2 := removeDigits(nums2, len(nums2)-cnt2)
		temp := make([]int, k)
		ti, i, j := 0, 0, 0
		for i < len(temp1) && j < len(temp2) {
			if maxArr(temp1[i:], temp2[j:]) {
				temp[ti] = temp1[i]
				i++
			} else {
				temp[ti] = temp2[j]
				j++
			}
			ti++
		}
		if i < len(temp1) {
			for i < len(temp1) {
				temp[ti] = temp1[i]
				i++
				ti++
			}
		}
		if j < len(temp2) {
			for j < len(temp2) {
				temp[ti] = temp2[j]
				j++
				ti++
			}
		}
		if maxArr(temp, ret) {
			ret = temp
		}
	}
	return ret
}

// 把数组当做整数看，a 是否大于等于 b
func maxArr(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return len(a) > len(b)
}

// 删除 k 个数字，使得剩余的数字组成的数字最大
// 实际上这就是 力扣 402、去除K位数字
func removeDigits(num []int, k int) []int {
	var stack []int
	removeCnt := 0
	for _, cur := range num {
		for len(stack) != 0 && cur > stack[len(stack)-1] && removeCnt < k {
			stack = stack[:len(stack)-1]
			removeCnt++
		}
		stack = append(stack, cur)
	}
	for removeCnt < k {
		stack = stack[:len(stack)-1]
		removeCnt++
	}
	for len(stack) > 1 && stack[0] == 0 {
		stack = stack[1:]
	}
	return stack
}
