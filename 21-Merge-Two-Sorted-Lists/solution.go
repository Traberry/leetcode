package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	}else {
		head = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node = l1
			l1 = l1.Next
		}else {
			node = l2
			l2 = l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node = l1
	}
	if l2 != nil {
		node = l2
	}

	return head
}

func main() {
	var a = []int{1, 3, 4}
	var b = []int{1, 2, 5}
	l := mergeTwoLists(itoLi(a), itoLi(b))
	fmt.Println(l.Next.Val)
}

//convert []int to ListNode
func itoLi(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode{Val:nums[0]}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val:nums[i]}
		temp = temp.Next
	}
	return res
}
