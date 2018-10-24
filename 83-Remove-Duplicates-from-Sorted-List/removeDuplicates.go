package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	node := head
	for node.Next != nil {
		N := node.Next
		if N.Val == node.Val {
			node.Next = N.Next
		}else {
			node = node.Next
		}
	}
	return head
}

func main() {
	var a = []int{1, 2, 3, 3, 4}
	l := ItoLi(a)
	res := deleteDuplicates(l)
	fmt.Println(res)
}

//transform []int to ListNode
func ItoLi(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var res = &ListNode{Val:nums[0]}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val:nums[i]}
		temp = temp.Next
	}
	return res
}

//transform ListNode to []int
func LitoI(l *ListNode) []int {
	if l == nil {
		return nil
	}

	var nums []int
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	return nums
}