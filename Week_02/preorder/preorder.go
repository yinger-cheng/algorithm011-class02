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
	node3 := &Node{3,[]*Node{node5,node6}}
	node2 := &Node{2,nil}
	node4 := &Node{4,nil}
	node1 := &Node{1,[]*Node{node3,node2,node4}}
	fmt.Println(preorder(node1))
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	result := []int{}
	result = preorderTraversal(root, result)
	return result
}

func preorderTraversal(root *Node,result []int)  []int {
	if root == nil {
		return result
	}
	result = append(result,root.Val)
	if root.Children != nil {
		for _ , node := range root.Children {
			childs := preorder(node)
			result = append(result,childs...)
		}
	}
	return result
}