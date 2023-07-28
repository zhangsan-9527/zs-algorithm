package double_linkedlist

import (
	"fmt"
	"testing"
)

func TestDoubleLinkedList_ToString(t *testing.T) {
	dlist := DoubleLinkedList{}
	dlist.Add(12)
	dlist.Add(22)
	dlist.Add(60)
	dlist.Add(80)
	dlist.Add(20)
	dlist.Add(99)
	dlist.Add(2)
	dlist.Add(9)
	dlist.Add(100)
	dlist.ToString()
	dlist.Delete(0)
	head := dlist.head
	fmt.Println(head)
	dlist.ToString()
}
