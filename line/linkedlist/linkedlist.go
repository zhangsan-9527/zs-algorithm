package linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	info interface{}
	next *Node
}

type LinkedList struct {

	// 长度
	len int

	// 头指针
	head *Node

	// 虚拟头节点(哨兵节点) 为了使头结点与其他节点拥有一样的处理逻辑
	//dummyHead *Node
}

func (l *LinkedList) GetLen() int {
	return l.len
}

// add 增加
func (l *LinkedList) Add(data interface{}) {
	newNode := &Node{info: data}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head // 时间复杂度O(N)
		// 不为nil 代表next有值, 然后就将next赋值给current进行下一个节点的查找
		for current.next != nil {
			current = current.next
		}
		current.next = newNode

	}
	l.len += 1
}

// insert 插入
func (l *LinkedList) Insert(index int, data interface{}) {
	// 1.创建一个node
	newNode := &Node{info: data}
	if l.len == 0 || index == 0 {
		newNode.next = l.head
		l.head = newNode
	} else {
		// 检查 index 的合法性
		if index < 0 || index >= l.len {
			panic("failed, index must >=0 or <len")
		}
		// 2.找到prev
		prev := l.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}
		// 3.将prev.next 传给新节点的next
		newNode.next = prev.next
		// 4. prev.next = newNode
		prev.next = newNode

	}
	l.len += 1

}

// get 查询指定索引的值
func (l *LinkedList) Get(index int) interface{} {
	current := &Node{}
	if index == 0 {
		current = l.head
	} else {
		// 检查 index 的合法性
		if index < 0 || index >= l.len {
			panic("failed, index must >=0 or <len")
		}

		current = l.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	}

	return current.info
}

// getFirst 获取第一个元素
func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

// getLast 获取最后一个元素
func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.len - 1)
}

// set 修改指定值index的val
func (l *LinkedList) Set(index int, data interface{}) {
	// 检查 index 的合法性
	if index < 0 || index >= l.len {
		panic("failed, index must >=0 or <len")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.info = data
}

// delete 删除指定节点
func (l *LinkedList) Delete(index int) interface{} {
	// 检查 index 的合法性
	if index < 0 || index >= l.len {
		panic("failed, index must >=0 or <len")
	}
	delNode := l.head
	if index == 0 {
		l.head = l.head.next
		delNode.next = nil
	} else {
		prev := l.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}
		//for i := 0; i < index; i++ {
		//	delNode = delNode.next
		//}
		delNode = prev.next
		prev.next = delNode.next
		delNode.next = nil
	}
	l.len -= 1
	return delNode.info
}

// contains 判断是否包含指定val
func (l *LinkedList) Contains(data interface{}) bool {
	current := l.head
	for current != nil {
		if current.info == data {
			return true
		}
		current = current.next
	}
	return false
}

// toString
func (l LinkedList) ToString() {
	sb := strings.Builder{}
	current := l.head
	for current != nil {
		strVal := fmt.Sprintf("%v => ", current.info)
		sb.WriteString(strVal)
		current = current.next
	}
	res := sb.String()
	fmt.Println(res)
}
