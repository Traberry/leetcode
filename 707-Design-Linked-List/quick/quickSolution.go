package quick

type MyLinkedList struct {
	List []int
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{List: make([]int, 0, 1000)}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= len(this.List) {
		return -1
	}

	return this.List[index]
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	var head = []int{val}
	this.List = append(head, this.List...)
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	this.List = append(this.List, val)
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index < 0 || index >= len(this.List) {
		return
	}

	this.List = append(this.List, 0)
	copy(this.List[index + 1:], this.List[index:])

	this.List[index] = val
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= len(this.List) {
		return
	}

	frontHalf := this.List[:index]
	backHalf := this.List[index + 1:]
	this.List = append(frontHalf, backHalf...)
}
