package double_linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	info interface{}
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	head *Node

	end *Node

	len int
}

// Len 双向链表长度
func (l DoubleLinkedList) Len() int {
	return l.len
}

// Add 增加元素
func (l *DoubleLinkedList) Add(data interface{}) {
	dNode := &Node{info: data}
	if l.head == nil {
		l.head = dNode
		l.end = dNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = dNode
		dNode.prev = current
		l.end = dNode

	}
	l.len += 1

}

// searchNode 查询指定index的节点
func (l *DoubleLinkedList) searchNode(index int) *Node {
	// 检查 index 的合法性
	if index < 0 || index >= l.len {
		return nil
	}
	if index < l.len/2 {
		current := l.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		return current
	} else {

		current := l.end
		for i := l.len - 1; i > index; i-- {
			current = current.prev
		}
		return current
	}
}

// Search 查询指定索引的元素值
func (l *DoubleLinkedList) Search(index int) interface{} {
	searchNode := l.searchNode(index)
	if searchNode == nil {
		panic("failed, index must >=0 or <len")
	}
	return searchNode.info
}

// SearchFirst 获取首元素
func (l *DoubleLinkedList) SearchFirst() interface{} {
	return l.head.info
}

// SearchEnd 获取尾元素
func (l *DoubleLinkedList) SearchEnd() interface{} {
	return l.head.info
}

// Update 修改指定索引的值
func (l *DoubleLinkedList) Update(index int, data interface{}) {

	upNode := l.searchNode(index)
	if upNode == nil {
		panic("failed, index must >=0 or <len")
	}
	upNode.info = data
}

// Insert 在指定位置插入元素
func (l *DoubleLinkedList) Insert(index int, data interface{}) {

	searchNode := l.searchNode(index)
	if searchNode == nil {
		panic("failed, index must >=0 or <len")
	}
	newNode := &Node{info: data}
	newNode.next = searchNode
	newNode.prev = searchNode.prev
	if searchNode.prev == nil {
		l.head = newNode
	} else {
		searchNode.prev.next = newNode
	}
	searchNode.prev = newNode
	l.len += 1

	// 通过if searchNode.prev == nil 可以直接判断是否是 index == 0
	//if index == 0 {
	//	l.head = newNode
	//}
	// 新节点 永远成为不了 end节点
	//if index == l.len-1 {
	//	l.end = newNode
	//}
}

// 删除指定节点
func (l *DoubleLinkedList) Delete(index int) interface{} {
	searchNode := l.searchNode(index)
	if searchNode == nil {
		panic("failed, index must >=0 or <len")
	}
	if searchNode.prev == nil {
		l.head = searchNode.next
	} else {
		searchNode.prev.next = searchNode.next
	}
	if searchNode.next == nil {
		l.end = searchNode.prev
	} else {
		searchNode.next.prev = searchNode.prev
	}
	l.len -= 1
	return searchNode.info

}

// ToString 字符串转换
func (l *DoubleLinkedList) ToString() {
	sb := strings.Builder{}
	current := l.head
	for current != nil {
		strVal := fmt.Sprintf("%v <=> ", current.info)
		sb.WriteString(strVal)
		current = current.next
	}
	res := sb.String()
	fmt.Println(res)
}
