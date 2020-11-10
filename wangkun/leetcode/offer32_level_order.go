package leetcode

/**
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]


提示：

节点总数 <= 1000
*/

func levelOrder(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root

	for len(queue) != 0 {
		r := queue[0]
		queue = queue[1:]
		ret = append(ret, r.Val)
		if r.Left != nil {
			queue = append(queue, r.Left)
		}
		if r.Right != nil {
			queue = append(queue, r.Right)
		}
	}

	return ret
}

/**
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

提示：

节点总数 <= 1000

*/
func levelOrder2(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	queue := []*TreeNode{root}

	level := 0
	for len(queue) != 0 {

		tmpQueue := make([]*TreeNode, 0)
		ret = append(ret, []int{})
		for _, v := range queue {
			ret[level] = append(ret[level], v.Val)
			if v.Left != nil {
				tmpQueue = append(tmpQueue, v.Left)
			}
			if v.Right != nil {
				tmpQueue = append(tmpQueue, v.Right)
			}
		}
		queue = tmpQueue
		level++
	}

	return ret
}

/**
请实现一个函数按照之字形顺序打印二叉树，即：
第一行按照从左到右的顺序打印，
第二层按照从右到左的顺序打印，
第三行再按照从左到右的顺序打印，其他行以此类推。

例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]

提示：

节点总数 <= 1000
*/
func levelOrder3(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	queue := []*TreeNode{root}

	level := 0
	for len(queue) != 0 {

		tmpQueue := make([]*TreeNode, 0)
		ret = append(ret, []int{})

		for _, v := range queue {
			ret[level] = append(ret[level], v.Val)
			if v.Left != nil {
				tmpQueue = append(tmpQueue, v.Left)
			}
			if v.Right != nil {
				tmpQueue = append(tmpQueue, v.Right)
			}
		}
		if level&1 == 1 {
			for i, j := 0, len(ret[level])-1; i < j; i, j = i+1, j-1 {
				ret[level][i], ret[level][j] = ret[level][j], ret[level][i]
			}
		}
		queue = tmpQueue
		level++
	}

	return ret
}

//还是第三题通过两个栈实现，也是官方的解法。
func levelOrder3ByStack(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}

	stack1 := []*TreeNode{root}
	var stack2 []*TreeNode

	level := 0
	for len(stack1) != 0 {

		ret = append(ret, []int{})
		stack2 = make([]*TreeNode, 0)
		for i := len(stack1) - 1; i >= 0; i-- {
			ret[level] = append(ret[level], stack1[i].Val)
			if stack1[i].Left != nil {
				stack2 = append(stack2, stack1[i].Left)
			}
			if stack1[i].Right != nil {
				stack2 = append(stack2, stack1[i].Right)
			}
		}

		level++
		if len(stack2) != 0 {
			ret = append(ret, []int{})
		}
		stack1 = make([]*TreeNode, 0)
		for i := len(stack2) - 1; i >= 0; i-- {
			ret[level] = append(ret[level], stack2[i].Val)
			if stack2[i].Right != nil {
				stack1 = append(stack1, stack2[i].Right)
			}
			if stack2[i].Left != nil {
				stack1 = append(stack1, stack2[i].Left)
			}
		}

		level++
	}

	return ret
}
