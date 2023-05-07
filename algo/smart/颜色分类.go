package smart

// 力扣 75、颜色分类
// 数组nums只包含0、1、2三个元素，用一次遍历对nums数组排序
// 这里的关键是 循环如何设计
func sortColors(nums []int) {
	h, t := 0, len(nums)-1
	for i := range nums {
		for ; i <= t && nums[i] == 2; t-- {
			nums[i], nums[t] = nums[t], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[h] = nums[h], nums[i]
			h++
		}
	}
}
