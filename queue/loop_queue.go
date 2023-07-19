package queue

import (
	"fmt"
	"strings"
)

type Loop_Queue struct {
	datas []interface{}
	head  int
	tail  int
	size  int
}

func NewLoop_Queue(cap int) *Loop_Queue {
	return &Loop_Queue{
		size:  0,
		datas: make([]interface{}, cap),
		head:  0,
		tail:  0,
	}
}

func (l *Loop_Queue) GetSize() int {
	return l.size
}

func (l *Loop_Queue) IsEmpty() bool {
	return l.head == l.tail
}

func (l *Loop_Queue) resize(str string) {
	var newList []interface{}
	if str == "up" {
		newList = make([]interface{}, 2*len(l.datas))
	} else {
		newList = make([]interface{}, len(l.datas)/2)
	}

	for i := 0; i < l.size; i++ {
		// newList[i] = l.datas[l.head]
		newList[i] = l.datas[(i+l.head)%len(l.datas)]
	}
	l.datas = newList
	l.head = 0
	l.tail = l.size

}

func (l *Loop_Queue) EnQueue(data interface{}) {
	if l.size == len(l.datas) {
		//panic("Queue 满了")
		l.resize("up")
	}
	l.datas[l.tail] = data
	l.size++
	l.tail = (l.tail + 1) % cap(l.datas)

}

func (l *Loop_Queue) DeQueue() interface{} {
	if l.size == 0 {
		panic("Queue 空的")
	}
	res := l.datas[l.head]
	l.datas[l.head] = nil
	l.size--
	l.head = (l.head + 1) % cap(l.datas)
	if l.size == cap(l.datas)/4 {
		l.resize("down")
	}
	return res
}

func (l *Loop_Queue) GetFront() interface{} {
	return l.datas[l.head]
}

func (l Loop_Queue) ToString() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Queue: size: %d, cap: %d\n", l.size, cap(l.datas)))
	sb.WriteString("队首: [")
	for i := 0; i < l.size; i++ {
		str := fmt.Sprintf("%v", l.datas[(i+l.head)%len(l.datas)])
		sb.WriteString(str)
		if (i+1)%len(l.datas) != l.tail {
			sb.WriteString(", ")
		}

	}
	sb.WriteString("]")
	return sb.String()

}
