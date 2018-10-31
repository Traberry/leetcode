package problem234

import (
	"testing"
	"errors"
)

func Test_isPalindrome2(t *testing.T) {
	for _, v := range tests {
		head := i2li(v.nums)
		if actual := isPalindrome2(head); actual != v.result {
			errors.New("test error")
		}
	}
}