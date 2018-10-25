package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		for head != nil {
			if head.Val == val {
				head = head.Next
			}else {
				break
			}
		}
	}

	node := head
	prev := node
	for node != nil {
		if node.Val == val {
			prev.Next = node.Next
		}else {
			prev = node
		}
		node = node.Next
	}
	return head
}

func main() {
	a := []int{1,2,3,3,4,5,6,6}
	test := ItoLi(a)
	fmt.Println(LitoI(removeElements(test, 3)))
}

func ItoLi(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var head = &ListNode{Val:nums[0]}
	node := head
	for i := 1; i < len(nums); i++ {
		node.Next = &ListNode{Val:nums[i]}
		node = node.Next
	}
	return head
}

func LitoI(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}