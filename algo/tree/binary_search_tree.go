package tree

import "math"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 深搜
// 力扣 98、验证二叉搜索树
// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树
// 有效 二叉搜索树定义如下：
// 		节点的左子树只包含 小于 当前节点的数
//		节点的右子树只包含 大于 当前节点的数
//		所有左子树和右子树自身必须也是二叉搜索树

func isValidBST(root *Node) bool {
	var stack []*Node
	pre := math.MinInt
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= pre {
			return false
		}
		pre = root.Val
		root = root.Right
	}
	return true
}
