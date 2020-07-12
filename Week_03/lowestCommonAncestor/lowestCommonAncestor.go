package main

import "fmt"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func main (){
	node7 := &TreeNode{7,nil,nil}
	node4 := &TreeNode{4,nil,nil}
	node2 := &TreeNode{2,node7,node4}
	node6 := &TreeNode{6,nil,nil}
	node5 := &TreeNode{5,node6,node2}

	node0 := &TreeNode{0,nil,nil}
	node8 := &TreeNode{8,nil,nil}
	node1 := &TreeNode{1,node0,node8}
	node3 := &TreeNode{3,node5,node1}

	node := lowestCommonAncestor(node3,node5,node1)
	fmt.Println(node.Val)
}

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	ltree := lowestCommonAncestor(root.Left, p, q)
	rtree := lowestCommonAncestor(root.Right, p , q)

	if ltree == nil {
		return rtree
	}else if rtree == nil {
		return ltree
	}else{
		return root
	}
}