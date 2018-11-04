package slice2stack

type MyStack struct {
	Nums []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{Nums: make([]int, 0, 1000)}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.Nums = append(this.Nums, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	p := this.Nums[len(this.Nums) - 1]
	this.Nums = this.Nums[:len(this.Nums)-1]
	return p
}


/** Get the top element. */
func (this *MyStack) Top() int {
	size := len(this.Nums)
	return this.Nums[size - 1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	size := len(this.Nums)
	return size == 0
}
