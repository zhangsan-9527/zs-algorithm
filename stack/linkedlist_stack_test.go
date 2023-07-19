package stack

import (
	"fmt"
	"testing"
)

func TestLinkedListStack_Peek(t *testing.T) {
	stack := NewLLStack()
	stack.Push("zhangsan")
	stack.Push("lisi")
	stack.Push("wangwu")
	stack.Push("zhaoliu")
	stack.datas.ToString()
	fmt.Println(stack.GetSize())
	res := stack.Peek()
	fmt.Println(res)
	stack.Pop()
	fmt.Println(stack.GetSize())
	res = stack.Peek()
	fmt.Println(res)
	stack.Pop()
	fmt.Println(stack.GetSize())
	res = stack.Peek()
	fmt.Println(res)
	fmt.Println(stack.GetSize())
	stack.Pop()
	res = stack.Peek()
	fmt.Println(res)
	fmt.Println(stack.GetSize())
}
