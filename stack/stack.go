package stack

type Stack interface {
	// GetSize 获取栈中元素个数
	GetSize() int
	// IsEmpty 判断栈中是否为空
	IsEmpty() bool
	// Push 向栈中放数据
	Push(data interface{})
	// Pop 从栈中弹出数据
	Pop() interface{}
	// Peek 查看栈顶元素
	Peek() interface{}
}
