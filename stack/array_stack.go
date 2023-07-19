package stack

import (
	"fmt"
	"strings"
)

var _ Stack = &ArrayStack{}

type ArrayStack struct {
	datas []interface{}
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{}
}

func (s *ArrayStack) GetSize() int {
	return len(s.datas)
}

func (s *ArrayStack) IsEmpty() bool {
	if len(s.datas) == 0 {
		return true
	} else {
		return false
	}
}

func (s *ArrayStack) Push(info interface{}) {
	s.datas = append(s.datas, info)
}

func (s *ArrayStack) Pop() interface{} {
	if len(s.datas) == 0 {
		panic("Stack is empty")
	}
	lastIndex := len(s.datas) - 1
	data := s.datas[lastIndex]
	s.datas = s.datas[:lastIndex]
	return data
}

func (s *ArrayStack) Peek() interface{} {
	if len(s.datas) == 0 {
		panic("Stack is empty")
	}
	data := s.datas[len(s.datas)-1]
	return data
}

func (s *ArrayStack) ToString() string {
	sb := strings.Builder{}
	sb.WriteString("Stack: [")
	for i := 0; i < len(s.datas); i++ {
		str := fmt.Sprintf("%v", s.datas[i])
		sb.WriteString(str)
		if i != len(s.datas)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
