package problem707

/****************** this method make problem more complex *********************/
/********** it's very hard to make it accepted by using this way *************/

//type MyLinkedList struct {
//	Val int
//	Next *MyLinkedList
//}
//
//
///** Initialize your data structure here. */
//func Constructor() MyLinkedList {
//	return MyLinkedList{Val:0, Next:nil}
//}
//
//
///** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
//func (this *MyLinkedList) Get(index int) int {
//	if this == nil {
//		return -1
//	}
//
//	temp := this
//	for i := 0; i < index; i++ {
//		temp = temp.Next
//		if temp == nil {
//			return -1
//		}
//	}
//	return temp.Val
//}
//
//
///** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
//func (this *MyLinkedList) AddAtHead(val int)  {
//	if this == nil {
//		return
//	}
//	var newHead = &MyLinkedList{Val: val}
//	newHead.Next = this
//	//this = newHead
//}
//
//
///** Append a node of value val to the last element of the linked list. */
//func (this *MyLinkedList) AddAtTail(val int)  {
//	if this == nil {
//		return
//	}
//	var newTail = &MyLinkedList{Val: val}
//	temp := this
//	for temp.Next != nil {
//		temp = temp.Next
//	}
//	temp.Next = newTail
//}
//
//
///** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
//func (this *MyLinkedList) AddAtIndex(index int, val int)  {
//	if this == nil{
//		return
//	}
//	var newNode = &MyLinkedList{Val: val}
//	temp := this
//	for i := 0; i < index-1; i++ {
//		temp = temp.Next
//		if temp == nil {
//			return
//		}
//	}
//	record := temp.Next
//	temp.Next = newNode
//	newNode.Next = record
//}
//
//
///** Delete the index-th node in the linked list, if the index is valid. */
//func (this *MyLinkedList) DeleteAtIndex(index int)  {
//	if this == nil {
//		return
//	}
//
//	if index == 0 {
//		this = this.Next
//		return
//	}
//
//	temp := this
//	for i := 0; i < index-1; i++ {
//		temp = temp.Next
//		if temp == nil {
//			return
//		}
//	}
//	record := temp.Next.Next
//	temp.Next = record
//}

