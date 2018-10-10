package stack

import "fmt"

type Stack []interface{}

func CreateStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value ...interface{}) {
	*s = append(*s, value...)
}

func (s *Stack) Pop() interface{} {
	p := (*s)[s.Size() - 1]
	*s = (*s)[:s.Size() - 1]
	return p
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Top() interface{} {
	if s.Size() > 0 {
		return (*s)[s.Size() - 1]
	}
	return nil
}

func (s *Stack) Empty() bool {
	if s.Size() > 0 {
		return false
	}else {
		return true
	}
}

//This Print() only fits in stack whose element is byte
func (s *Stack) Print() {
	for _, v := range *s {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
}
