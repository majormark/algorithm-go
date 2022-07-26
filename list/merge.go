package list

/**
合并两个有序链表 p.84
*/
func Merge(h1 *Node, h2 *Node) *Node {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	n1 := h1
	n2 := h2
	var newHead *Node
	var newTail *Node
	for n1 != nil && n2 != nil {
		var cur *Node
		if n1.Val < n2.Val {
			cur = n1
			n1 = n1.Next
		} else {
			cur = n2
			n2 = n2.Next
		}
		if newHead == nil {
			newHead = cur
			newTail = cur
		} else {
			newTail.Next = cur
		}
	}
	for n1 != nil {
		newTail.Next = n1
		n1 = n1.Next
	}
	for n2 != nil {
		newTail.Next = n2
		n2 = n2.Next
	}
	return newHead
}
