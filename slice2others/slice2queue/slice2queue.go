package slice2queue

type Queue struct {
	nums []int
}

func CreateQueue() *Queue {
	return &Queue{nums: []int{}}
}

func (q *Queue) Push(v int) {
	q.nums = append(q.nums, v)
}

func (q *Queue) Pop() int {
	return q.nums[0]
}

func (q *Queue) length() int {
	return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
	return q.length() == 0
}