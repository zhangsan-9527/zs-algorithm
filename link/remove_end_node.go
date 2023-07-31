package link

import (
	"fmt"
	"strings"
)

/*
	给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
	给定的 n 保证是有效的。
	你能尝试使用一趟扫描实现吗？


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(fst int) *ListNode {
	return &ListNode{
		Val: fst,
	}
}

func (l *ListNode) AddNode(val int) {
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}

	current := l
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

}

// RemoveNthFromEnd 链表获取倒数第N位的值
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	var pre *ListNode
	cur := result
	i := 1
	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return result.Next
}

func (l ListNode) ToString() {
	sb := strings.Builder{}
	current := l
	for current.Next != nil {
		strVal := fmt.Sprintf("%v => ", current.Val)
		sb.WriteString(strVal)
		current = *current.Next
	}
	strVal := fmt.Sprintf("%v", current.Val)
	sb.WriteString(strVal)
	res := sb.String()
	fmt.Println(res)
}
