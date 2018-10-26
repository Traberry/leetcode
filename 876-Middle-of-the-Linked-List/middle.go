package Middle_of_the_Linked_List


//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 0
	node := head
	for node != nil {
		length++
		node = node.Next
	}

	for i := 0; i < length / 2; i++ {
		head = head.Next
	}

	return head


}