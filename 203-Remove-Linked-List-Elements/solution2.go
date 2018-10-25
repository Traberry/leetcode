package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements2(head *ListNode, val int) *ListNode {
	headPre := &ListNode{Next:head}

	temp := headPre
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		}else {
			temp = temp.Next
		}
	}
	return headPre.Next
}

func main() {
	a := []int{1,2,3,3,4,5,6,6}
	test := ItoLi(a)
	fmt.Println(LitoI(removeElements2(test, 3)))
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
