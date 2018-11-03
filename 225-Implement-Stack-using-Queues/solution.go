package problem225

type MyStack struct {
	Queue []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{Queue: make([]int, 0, 1000)}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.Queue = append(this.Queue, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	p := this.Queue[len(this.Queue) - 1]
	this.Queue = this.Queue[:len(this.Queue)-1]
	return p
}


/** Get the top element. */
func (this *MyStack) Top() int {
	size := len(this.Queue)
	return this.Queue[size - 1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	size := len(this.Queue)
	return size == 0
}
