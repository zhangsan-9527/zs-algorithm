package linkedlista

import (
	"fmt"
	"testing"
)

func TestLinkedList_delete(t *testing.T) {
	linklist := LinkedList[int]{}
	linklist.add(12)
	linklist.add(15)
	linklist.add(17)
	linklist.add(7)
	linklist.add(50)
	linklist.add(60)
	linklist.add(70)

	res := linklist.contains(80)
	fmt.Println(res)
	linklist.toString()
}
