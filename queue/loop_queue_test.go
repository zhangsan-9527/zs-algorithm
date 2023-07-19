package queue

import (
	"fmt"
	"testing"
)

func TestLoop_Queue_GetSize(t *testing.T) {
	queue := NewLoop_Queue(5)
	queue.EnQueue(45)
	queue.EnQueue(55)
	queue.EnQueue(82)
	queue.EnQueue(64)
	queue.EnQueue(88)
	queue.EnQueue(88)
	queue.EnQueue(88)
	queue.EnQueue(88)
	str := queue.ToString()
	fmt.Println(str)
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	str = queue.ToString()
	fmt.Println(str)

}
