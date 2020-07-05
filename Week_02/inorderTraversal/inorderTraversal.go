package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	//创建一颗树
	node3 := &TreeNode{3,nil,nil}
	node2 := &TreeNode{2,node3,nil}
	node1 := &TreeNode{1,nil,node2}
	//递归
	fmt.Println(inorderTraversal(node1))
	//栈
	fmt.Println(inorderTraversal1(node1))
}
/**
 * 递归算法
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	data := append(inorderTraversal(root.Left),root.Val)
	data = append(data,inorderTraversal(root.Right)...)
	return data
}

/**
 * 迭代，使用栈存左节点
 */
func inorderTraversal1(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1
		root = stack[index]
		result = append(result, root.Val)
		root = root.Right
		stack = stack[:index]
	}
	return result
}

func inorderTraversal2(root *TreeNode) []int{
	result := []int{}
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		index := len(stack) - 1
		root = stack[index]
		result = append(result,root.Val)
		root = root.Right
		stack = stack[:index]
	}
	return result

}


























