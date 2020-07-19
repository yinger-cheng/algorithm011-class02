package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main(){
	//创建一颗树
	//node35 := &TreeNode{5,nil,nil}
	//node33 := &TreeNode{3,nil,nil}
	//node39 := &TreeNode{9,nil,nil}
	//node23 := &TreeNode{3,node35,node33}
	//node22 := &TreeNode{2,nil,node39}
	//node1 := &TreeNode{1,node23,node22}
	//fmt.Println(largestValues(node1))

	node21 := &TreeNode{-1,nil,nil}
	node1 := &TreeNode{0,node21,nil}
	fmt.Println(largestValues(node1))
}
//层序遍历
func largestValues(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return  result
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0 ; i++ {
		tmp := []*TreeNode{}
		max := queue[0].Val
		for j := 0; j < len(queue) ; j++  {
			node := queue[j]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				tmp = append(tmp,node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp,node.Right)
			}
		}
		result = append(result,max)
		queue = tmp
	}
	return result
}
