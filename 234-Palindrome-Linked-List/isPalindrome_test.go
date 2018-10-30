package problem234

import (
	"testing"
	"errors"
)

var tests = []struct {
	nums []int
	result bool
}{
	{[]int{1,1,2,2,3,3}, false},
	{[]int{1,2,2,1}, true},
}

func Test_isPalindrome(t *testing.T) {
	for _, test := range tests {
		head := i2li(test.nums)
		if actual := isPalindrome(head); actual != test.result {
			errors.New("test error")
		}
	}
}

func i2li(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var head = &ListNode{Val:nums[0]}
	temp := head
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val:nums[i]}
		temp = temp.Next
	}
	return head
}