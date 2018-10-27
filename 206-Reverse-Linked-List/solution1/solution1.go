package solution1


// Definition for singly-linked list.
type listNode struct {
	Val int
	Next *listNode
}

func reverseList(head *listNode) *listNode {
	oriNums := litoI(head)
	revNums := numsReverse(oriNums)
	return itoLi(revNums)
}

func litoI(head *listNode) []int {
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

func itoLi(nums []int) *listNode {
	if len(nums) == 0 {
		return nil
	}

	var head = &listNode{Val: nums[0]}
	node := head
	for i := 1; i < len(nums); i++ {
		node.Next = &listNode{Val: nums[i]}
		node = node.Next
	}
	return head
}

func numsReverse(nums []int) []int {
	var reNums []int
	for i := len(nums)-1; i >= 0; i-- {
		reNums = append(reNums, nums[i])
	}
	return reNums
}
