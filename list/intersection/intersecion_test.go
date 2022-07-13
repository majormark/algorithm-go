package intersection

import (
	"fmt"
	"testing"
)

func BuildIntersectionList() *Node {
	head := NewNode(0)
	tmp := head
	var mid *Node
	for i := 1; i < 7; i++ {
		tmp.Next = NewNode(i)
		tmp = tmp.Next
		if i == 4 {
			mid = tmp
		}
	}
	tmp.Next = mid
	return head
}
func TestGetLoopNode(t *testing.T) {
	h := BuildIntersectionList()
	n := GetLoopNode(h)
	if n == nil {

		fmt.Println(fmt.Sprintf("TestGetLoopNode has no loopNode"))
	} else {

		fmt.Println(fmt.Sprintf("TestGetLoopNode loopNode: %v", n.Val))
	}
}
