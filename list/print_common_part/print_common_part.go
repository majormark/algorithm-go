package print_common_part

import "fmt"

/**
打印两个有序链表的公共部分
*/
type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

func PrintCommonPart(h1 *Node, h2 *Node) {
	if h1 == nil || h2 == nil {
		return
	}
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			h1 = h1.Next
		} else if h1.Val > h2.Val {
			h2 = h2.Next
		} else {
			fmt.Println(h1.Val)
		}
	}
}
