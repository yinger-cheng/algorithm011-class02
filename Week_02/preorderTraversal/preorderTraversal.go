package main

import "fmt"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func main(){
	//创建一颗树
	node3 := &TreeNode{3,nil,nil}
	node2 := &TreeNode{2,node3,nil}
	node1 := &TreeNode{1,nil,node2}
	fmt.Println(preorderTraversal(node1))
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
	 	return result
	 }
	 result = append(result,root.Val)
	 if root.Left != nil {
	 	 tmp := preorderTraversal(root.Left)
		 result = append(result,tmp...)
	 }
	if root.Right != nil {
		tmp := preorderTraversal(root.Right)
		result = append(result,tmp...)
	}
	return result
}