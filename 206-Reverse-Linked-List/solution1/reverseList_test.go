package solution1

import "testing"

type sample struct {
	inPut *listNode
	outPut *listNode
}

func TestReverseList(t *testing.T) {
	a1 := itoLi([]int{1, 2, 3, 4, 5})
	a2 := itoLi([]int{5, 4, 3, 2, 1})
	b1 := itoLi([]int{7, 8, 9})
	b2 := itoLi([]int{9, 8, 7})

	tests := []sample{
		{a1, a2},
		{b1, b2},
	}

	for _, test := range tests {
		if actual := reverseList(test.inPut); actual != test.outPut { // variable "actual" and "test.outPut" both are pointer(address), they cannot be equal,
			t.Error("error")                                     // so how to test pointer ???
		}
	}

}
