package tree

import "fmt"

func preOrder(root *Node) (res []int) {
	if root != nil {
		stack := []*Node{root}
		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}
	return res
}

func inOrder(root *Node) (res []int) {
	if root != nil {
		var stack []*Node
		cur := root
		for cur != nil || len(stack) > 0 {
			for cur != nil {
				stack = append(stack, cur)
				cur = cur.Left
			}
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

func postOrder(root *Node) (res []int) {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{cur.Val}, res...)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}
	return res
}

func test() {
	//      1
	//     / \
	//    2   3
	// 	 / \   \
	//	4   5   6
	node6 := &Node{6, nil, nil}

	node4 := &Node{4, nil, nil}
	node5 := &Node{5, nil, nil}

	node3 := &Node{3, nil, node6}
	node2 := &Node{2, node4, node5}
	root := &Node{1, node2, node3}

	fmt.Printf("先序遍历结果为：%v\n", preOrder(root))  // [1 2 4 5 3 6]
	fmt.Printf("中序遍历结果为：%v\n", inOrder(root))   // [1 2 4 5 3 6]
	fmt.Printf("后序遍历结果为：%v\n", postOrder(root)) // [1 2 4 5 3 6]
}
