package p155_quick

type elem struct {
	val int
	min int
}

type MinStack struct {
	stack []elem
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	m := x
	if len(this.stack) > 0 && this.GetMin() < x {
		m = this.GetMin()
	}
	this.stack = append(this.stack, elem{val: x, min: m})
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack) - 1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1].val
}


func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack) - 1].min
}
