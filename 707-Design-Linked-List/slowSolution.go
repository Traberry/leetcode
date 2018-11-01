package problem707

// MyLinkedList用来记录整个单链表的状态
type MyLinkedList struct {
	size int
	head, tail *ListNode
}

// ListNode是链表中每个节点的结构体
type ListNode struct {
	Val int
	Next *ListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	// 创建一个头节点，指向一个尾节点（不是nil）
	t := &ListNode{}
	h := &ListNode{Next: t}
	return MyLinkedList{
		head: h,
		tail: t,
	}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	temp := this.head.Next
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	return temp.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	// 构建一个节点来承载值
	node := &ListNode{Val: val}
	// 构建的节点指向头节点的下一个节点
	node.Next = this.head.Next
	// 头节点不变，指向将要加入的节点
	this.head.Next = node

	this.size++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	// 把值赋给当前的尾节点
	this.tail.Val = val

	// 生成一个新的、空的尾节点
	node := &ListNode{}
	this.tail.Next = node
	this.tail = node

	this.size++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.size || index < 0 {
		return
	}else if index == this.size {
		this.AddAtTail(val)

		this.size++
	}else {
		temp := this.head.Next
		for i := 0; i < index-1; i++ {
			temp = temp.Next
		}
		node := &ListNode{Val: val}
		node.Next = temp.Next
		temp.Next = node

		this.size++
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.size {
		return
	}

	temp := this.head.Next
	for i := 0; i < index-1; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next

	this.size--

}
