package main

import (
	"fmt"
	"leetcode/20-Valid-Parentheses/stack"
)

func isValid(s string) bool {
	m := make(map[byte]byte)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'

	testStack := stack.CreateStack()
	testStack.Push(s[0])
	
	for i := 1; i < len(s); i++ {
		if m[ s[i] ] == testStack.Top() {
			testStack.Pop()
		}else {
			testStack.Push(s[i])
		}
	}

	return testStack.Empty()
}

func main() {
	str := "()()(){[()]}"
	fmt.Println(isValid(str))
}
