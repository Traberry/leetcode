package realSolution

import "leetcode/slice2others/slice2queue"

type MyStack struct {
	a, b *slice2queue.Queue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{a: slice2queue.CreateQueue(), b: slice2queue.CreateQueue()}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.a.Push(x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.a.Length() == 0 {
		this.a, this.b = this.b, this.a
	}

	for this.a.Length() > 1 {
		this.b.Push(this.a.Pop())
	}

	this.a.Pop()
}


/** Get the top element. */
func (this *MyStack) Top() int {
	t := this.Pop()
	this.Push(t)
	return t
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.a.Length() == 0 && this.b.Length() == 0
}
