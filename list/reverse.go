package list

func ReverseSingle(h *Node) *Node {
	if h == nil {
		return nil
	}
	if h.Next == nil {
		return h
	}
	cur := h
	var pre *Node
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func ReverseDouble(h *DoubleNode) *DoubleNode {
	if h == nil {
		return nil
	}
	if h.Next == nil {
		return h
	}
	cur := h
	var pre *DoubleNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		cur.Last = next
		pre = cur
		cur = next
	}

	return pre
}

func ReversePart(h *Node, from int, to int) *Node {
	if h == nil {
		return nil
	}
	cur := h
	var tPre *Node
	var tPos *Node
	var length int
	for cur != nil {
		length++
		if length == from-1 {
			tPre = cur
		}
		if length == to+1 {
			tPos = cur
		}
		cur = cur.Next
	}
	if from > to || from < 1 || to > length {
		return h
	}
	node1 := h
	if tPre != nil {
		node1 = tPre.Next
	}
	node2 := node1.Next
	node1.Next = tPos
	for node2 != tPos {
		next := node2.Next
		node2.Next = node1
		node1 = node2
		node2 = next
	}
	if tPre == nil {
		return node1
	}
	tPre.Next = node1
	return h

}
