package problem232

type MyQueue struct {
	a, b *stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{a: newStack(), b: newStack()}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.a.push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.b.length() == 0 {
		for this.a.length() > 0 {
			this.b.push(this.a.pop())
		}
	}
	return this.b.pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	p := this.Pop()
	this.b.push(p)
	return p
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.a.isEmpty() && this.b.isEmpty()
}


/********************* implement of stack **************************/
type stack struct {
	nums []int
}

func newStack() *stack {
	return &stack{nums: []int{}}
}

func (s *stack) push(v int) {
	s.nums = append(s.nums, v)
}

func (s *stack) pop() int {
	p := s.nums[len(s.nums) - 1]
	s.nums = s.nums[:len(s.nums) - 1]
	return p
}

func (s *stack) length() int {
	return len(s.nums)
}

func (s *stack) isEmpty() bool {
	return len(s.nums) == 0
}
