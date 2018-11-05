package problem844

func backspaceCompare(S string, T string) bool {
	// 构造两个栈，把字符串S和T中的每个字母都入栈
	s := Constructor()
	t := Constructor()
	for i := 0; i < len(S); i++ {
		s.Push(S[i])
	}
	for i := 0; i < len(T); i++ {
		t.Push(T[i])
	}

	sBytes := out(&s)
	tBytes := out(&t)

	if len(sBytes) != len(tBytes) {
		return false
	}

	for i := 0; i < len(sBytes); i++ {
		if sBytes[i] != tBytes[i] {
			return false
		}
	}

	return true
}

// 按照题目规则（# 代表删除键）把字符串中的字母出栈
func out(stack *MyStack) []byte {
	var out []byte
	var count = 0
	for !stack.Empty() {
		b := stack.Pop()
		if b == '#' {
			count++
		}else {
			if count > 0 {
				count--
			}else {
				out = append(out, b)
			}
		}
	}
	return out
}


/**********************byte stack**************************/
type MyStack struct {
	Bytes []byte
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{Bytes: []byte{}}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x byte)  {
	this.Bytes = append(this.Bytes, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() byte {
	p := this.Bytes[len(this.Bytes) - 1]
	this.Bytes = this.Bytes[:len(this.Bytes)-1]
	return p
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	size := len(this.Bytes)
	return size == 0
}