package list

/**
向有序环形单链表插入节点，p.82
*/
func insertNum(h *Node, num int) *Node {
	n := &Node{Val: num}
	if h == nil {
		n.Next = n
		h = n
		return h
	}
	pre := h
	cur := h.Next
	for cur != h {
		if pre.Val <= num && num <= cur.Val {
			break
		}
		pre = cur
		cur = cur.Next
	}
	pre.Next = n
	n.Next = cur
	if n.Val < h.Val {
		h = n
	}
	return h
}
