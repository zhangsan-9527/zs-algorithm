package stack

import (
	"fmt"
	"testing"
)

func TestArrayStack_Push(t *testing.T) {
	stack := NewArrayStack()
	stack.Push("zhangsan")
	stack.Push("lisi")
	stack.Push("wangwu")
	stack.Push("zhaoliu")
	res1 := stack.ToString()
	fmt.Println(res1)
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
