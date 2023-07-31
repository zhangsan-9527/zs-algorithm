package link

import (
	"testing"
)

func TestRemoveEndNode(t *testing.T) {
	list := NewListNode(1)

	list.AddNode(99)
	list.AddNode(80)
	list.AddNode(800)

	list.ToString()
	//RemoveNthFromEnd(&list, 2)
	//list.ToString()
}
