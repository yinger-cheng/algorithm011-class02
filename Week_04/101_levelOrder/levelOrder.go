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
	node15 := &TreeNode{15,nil,nil}
	node7 := &TreeNode{7,nil,nil}
	node20 := &TreeNode{20,node15,node7}
	node9 := &TreeNode{9,nil,nil}
	node3 := &TreeNode{3,node9,node20}
	fmt.Println(levelOrder1(node3))
}
//层序遍历
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return  result
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0 ; i++ {
		result = append(result,[]int{})
		tmp := []*TreeNode{}
		for j := 0; j < len(queue) ; j++  {
			node := queue[j]
			result[i] = append(result[i],node.Val)
			if node.Left != nil {
				tmp = append(tmp,node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp,node.Right)
			}
		}
		queue = tmp
	}
	return result
}

//dfs实现
func levelOrder1(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return  nil
	}
	result = dfsLevelOrder1(0,root,result)
	return result
}

func dfsLevelOrder1(level int,root *TreeNode,result [][]int)  [][]int{
	if root == nil {
		return result
	}
	if level == len(result){
		result = append(result,[]int{})
	}
	result[level] = append(result[level],root.Val)
	result = dfsLevelOrder1(level+1,root.Left,result)
	result = dfsLevelOrder1(level+1,root.Right,result)
	return result
}


























