package link

import (
	"testing"
	"zs-alogrithm/link/linkedlist"
)

func TestRemoveEndNode(t *testing.T) {
	list := linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)
	list.ToString()
	list.Insert(5, 999)
	list.ToString()
}
