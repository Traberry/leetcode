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
	p := q.nums[0]
	q.nums = q.nums[1:]
	return p
}

func (q *Queue) Length() int {
	return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}