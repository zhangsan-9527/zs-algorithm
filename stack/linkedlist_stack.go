package stack

import (
	"algo/line/linkedlist"
)

type LinkedListStack struct {
	datas linkedlist.LinkedList
}

func NewLLStack() *LinkedListStack {
	return &LinkedListStack{}
}

func (l *LinkedListStack) GetSize() int {
	return l.datas.GetLen()
}

func (l *LinkedListStack) IsEmpty() bool {
	if l.datas.GetLen() == 0 {
		return true
	} else {
		return false
	}
}

func (l *LinkedListStack) Push(data interface{}) {
	l.datas.Insert(0, data)
}

func (l *LinkedListStack) Pop() interface{} {
	if l.datas.GetLen() == 0 {
		panic("Stack is empty")
	}
	data := l.datas.Delete(0)
	return data
}

func (l *LinkedListStack) Peek() interface{} {
	data := l.datas.GetFirst()
	return data
}
