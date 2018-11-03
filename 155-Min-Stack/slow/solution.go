package p155_slow

type MinStack struct {
	Queue []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Queue: make([]int, 0, 1000)}
}


func (this *MinStack) Push(x int)  {
	this.Queue = append(this.Queue, x)
}


func (this *MinStack) Pop()  {
	this.Queue = this.Queue[:len(this.Queue)-1]
}


func (this *MinStack) Top() int {
	size := len(this.Queue)
	return this.Queue[size - 1]
}


func (this *MinStack) GetMin() int {
	min := this.Queue[0]
	for i := 1; i < len(this.Queue); i++ {
		if this.Queue[i] < min {
			min = this.Queue[i]
		}
	}
	return min
}
