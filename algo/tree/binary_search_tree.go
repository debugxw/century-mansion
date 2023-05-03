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

// 思路非常不错
// 力扣 538、将二叉搜索树转换为累加树
// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree）
// 使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和
func convertBST(root *Node) *Node {
	sum := 0
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root != nil {
			// 反序遍历的思路很好
			dfs(root.Right)
			sum += root.Val
			root.Val = sum
			dfs(root.Left)
		}
	}
	dfs(root)
	return root
}
