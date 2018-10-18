package main

import "testing"

type Buddy struct {
	s1 string
	s2 string
	isBuddy bool
}

func TestBuddyStrings(t *testing.T) {
	tests := []Buddy {
		{"aab", "aac", true},
		{"abcd", "abcd", false},
		{"ab", "ba", true},
		{"a", "aa", false},
		{"caaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", "caaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
	}

	for _, test := range tests {
		if actual := buddyStrings(test.s1, test.s1); actual != test.isBuddy {
			t.Errorf("buddyStrings(%s, %s); Got %t; Expected %t;", test.s1, test.s2, actual, test.isBuddy)
			//t.Error("error")
		}
	}
}
