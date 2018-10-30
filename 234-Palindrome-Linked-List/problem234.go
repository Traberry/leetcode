package problem234

type ListNode struct {
	Val int
	Next *ListNode
}

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
