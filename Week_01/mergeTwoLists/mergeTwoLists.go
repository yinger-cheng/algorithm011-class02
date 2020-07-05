package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func main() {
	//mylist := list.New()
	//mylist.PushFront(2)
	//mylist.PushBack(1)


	l1_3 := ListNode{
		Val:4,
		Next:nil,
	}
	l1_2 := ListNode{
		2,
		&l1_3,
	}
	l1_1 := ListNode{
		1,
		&l1_2,
	}

	l2_3 := ListNode{
		Val:4,
		Next:nil,
	}
	l2_2 := ListNode{
		3,
		&l2_3,
	}
	l2_1 := ListNode{
		1,
		&l2_2,
	}
	
	l3 := mergeTwoLists(&l1_1,&l2_1)

	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	minNode := &ListNode{}
	result := minNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			minNode.Next = l1
			l1 = l1.Next
		}else{
			minNode.Next = l2
			l2 = l2.Next
		}
		minNode = minNode.Next
	}
	if l1 != nil {
		minNode.Next = l1
	}
	if l2 != nil {
		minNode.Next = l2
	}
	return result.Next
}