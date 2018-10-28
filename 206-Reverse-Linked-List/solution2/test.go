package main

type listNode struct {
	Val int
	Next *listNode
}

// 总体思路：把链表中的节点一个一个地逆转
func reverseListNodePrint(head *listNode) *listNode {
	// 创建一个空节点，用来被head指向
	var prev *listNode

	for head != nil {
		// 提前记录下一个要被逆转的节点
		temp := head.Next

		// 把当前节点的Next指向prev
		head.Next = prev

		// 让当前的节点成为prev
		prev = head

		// 让head指向下一个要被逆转的节点
		head = temp
	}

	return prev
}

func main() {
	
}
