package main

import (
	"strings"
	"fmt"
)

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if strings.Compare(A, B) == 0 {
		var count [26]int
		for i := 0; i < len(A); i++ {
			count[ A[i] - 'a' ]++
		}
		for _, v := range count {
			if v > 1 {
				return true
			}
		}
		return false
	}else {
		first, second := -1, -1
		for i := 0; i < len(A); i++ {
			if A[i] != B[i] {
				if first == -1 {
					first = i
				}else if second == -1 {
					second = i
				}else {
					return false
				}
			}
		}

		return A[first] == B[second] && A[second] == B[first]
	}
}

func main() {
	A := "ab"
	B := "ba"
	fmt.Println(buddyStrings(A, B))
}
