package queue

type Queue interface {

	// GetSize 查询队列中元素个数
	GetSize() int

	// IsEmpty 判断是否为空
	IsEmpty() bool

	// EnQueue 入队
	EnQueue(data interface{})

	// DeQueue 出队
	DeQueue() interface{}

	// GetFront 查看队首元素
	GetFront() interface{}
}
