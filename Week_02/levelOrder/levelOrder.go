package main

import "fmt"

type Node struct {
	Val int
	Children []*Node
}

func main(){
	//创建一颗树
	node5 := &Node{5,nil}
	node6 := &Node{6,nil}
	node7 := &Node{7,nil}
	node3 := &Node{3,[]*Node{node5,node6}}
	node2 := &Node{2,[]*Node{node7}}
	node4 := &Node{4,nil}
	node1 := &Node{1,[]*Node{node3,node2,node4}}
	fmt.Println(levelOrder(node1))
}

func levelOrder(root *Node)  [][]int {
	result := [][]int{}
	result = levelOrderTraversal(root,0, result)
	return result
}

func levelOrderTraversal(root *Node,level int,result [][]int) [][]int {
	if root == nil {
		return result
	}
	if len(result) == level {
		result = append(result, []int{})
	}
	result[level] = append(result[level],root.Val)
	for _ , node := range root.Children {
		result = levelOrderTraversal(node,level+1,result)
	}
	return result
}
