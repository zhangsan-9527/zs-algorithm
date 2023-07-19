package linkedlista

import (
	"fmt"
	"reflect"
	"strings"
)

type Node[T any] struct {
	info T
	next *Node[T]
}

type LinkedList[T any] struct {

	// 长度
	len int

	// 头指针
	head *Node[T]

	// 虚拟头节点(哨兵节点) 为了使头结点与其他节点拥有一样的处理逻辑
	//dummyHead *Node
}

func NewList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add 增加
func (l *LinkedList[T]) add(data T) {
	newNode := &Node[T]{info: data}
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
func (l *LinkedList[T]) insert(index int, data T) {
	// 检查 index 的合法性
	if index < 0 || index >= l.len {
		panic("failed, index must >=0 or <len")
	}
	// 1.创建一个node
	newNode := &Node[T]{info: data}
	if l.len == 0 {
		l.head = newNode
	} else {
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

}

// get 查询指定索引的值
func (l *LinkedList[T]) get(index int) T {
	// 检查 index 的合法性
	if index < 0 || index >= l.len {
		panic("failed, index must >=0 or <len")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.info
}

// getFirst 获取第一个元素
func (l *LinkedList[T]) getFirst() T {
	return l.get(0)
}

// getLast 获取最后一个元素
func (l *LinkedList[T]) getLast() T {
	return l.get(l.len - 1)
}

// set 修改指定值index的val
func (l *LinkedList[T]) set(index int, data T) {
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
func (l *LinkedList[T]) delete(index int) T {
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
	return delNode.info
}

// contains 判断是否包含指定val
func (l *LinkedList[T]) contains(data T) bool {
	current := l.head
	for current != nil {
		if reflect.DeepEqual(current.info, data) {
			//if cmp.Equal(current.info, data) {
			return true
		}
		current = current.next
	}
	return false
}

// toString
func (l LinkedList[T]) toString() {
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
