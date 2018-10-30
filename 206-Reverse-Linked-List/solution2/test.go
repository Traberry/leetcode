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

// 思路与上面的方法相同，不过把主要操作封装成一个递归的函数
func reverseListNode2(head *listNode) *listNode {
	// 一开始传入头节点和一个nil
	return recursively(head, nil)
}

// 传入将要被逆转的节点cur以及上一个被逆转的节点prev
func recursively(cur, prev *listNode) *listNode {
	// 操作进行到最后一个节点，结束
	if cur == nil {
		return prev
	}

	// 提前记录下一个要被逆转的节点
	next := cur.Next

	// 把当前节点的Next指向prev
	cur.Next = prev

	return recursively(next, cur)
}