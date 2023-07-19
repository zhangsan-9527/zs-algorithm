package queue

import "zs-alogrithm/line/linkedlist"

type Linked_Quene struct {
	datas linkedlist.LinkedList
}

func (l *Linked_Quene) GetSize() int {
	return l.datas.GetLen()
}

func (l *Linked_Quene) IsEmpty() bool {
	if l.datas.GetLen() == 0 {
		return true
	} else {
		return false
	}
}

func (l *Linked_Quene) EnQueue(data interface{}) {
	l.datas.Insert(0, data)
}

func (l *Linked_Quene) DeQueue() interface{} {
	if l.datas.GetLen() == 0 {
		panic("Queue is empty")
	}
	dleData := l.datas.Delete(l.datas.GetLen() - 1)
	return dleData
}

func (l *Linked_Quene) GetFront() interface{} {
	return l.datas.GetFirst()
}
