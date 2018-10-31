package problem234

type ListNode struct {
	Val int
	Next *ListNode
}

/*************************** solution1 ****************************/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	nums := litoi(head)
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func litoi(head *ListNode) []int {
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

/*************************** solution2 ****************************/
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// let slow point to the middle of the linked list
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next
	}

	// reverse slow: reverse the right half of the linked list
	slow = reverse(slow)

	// compare the letf half to the reversed right half
	comp := head
	for slow != nil {
		if slow.Val != comp.Val {
			return false
		}
		slow = slow.Next
		comp = comp.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}