package main

import (
	"fmt"
	"math/rand"
)

type tree struct {
	data int
	left *tree
	right *tree
}

// 创建二叉树
func createTree(arr []int) []tree {
	d := make([]int, 0)
	st := &tree{}
	for i, ar := range arr {
		d = append(d, st)
		d[i].data = ar
	}

	for i := 0; i < len(arr)/2; i++ {
		d[i].left = &d[i*2+1]
		if i*2+2 < len(d) {
			d[i].right = &d[i*2+2]
		}
	}
	fmt.Println(d)
	return d
}

// 前序遍历
func preorder(root tree) {
	fmt.Println(root.data, " ")
	if root.left != nil {
		preorder(*root.left)
	}
	if root.right != nil {
		preorder(*root.right)
	}
}

// 中序遍历
func inorder(root tree) {
	if root.left != nil {
		inorder(*root.left)
	}
	fmt.Println(root.data," ")
	if root.right != nil {
		inorder(*root.right)
	}
}

// 后序遍历
func afterorder(root tree) {
	if root.left != nil {
		afterorder(*root.left)
	}
	if root.right != nil {
		afterorder(*root.right)
	}
	fmt.Println(root.data, " ")
}

func main() {
	arr := make([]int, 0)
	for i:=0; i < 20; i++ {
		arr = append(arr, rand.Intn(80))
	}

	fmt.Println(arr)
	t := createTree(arr)
	preorder(t[0])
	fmt.Println()
	inorder(t[0])
	fmt.Println()
	afterorder(t[0])
}
