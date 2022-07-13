package remove_last_kth_node

type Node struct {
	Val  int
	Next *Node
}

type DoubleNode struct {
	Val  int
	Next *DoubleNode
	Last *DoubleNode
}

func RemoveLastKthNodeSingle(h *Node, lastK int) *Node {
	if h == nil {
		return nil
	}
	cur := h
	for cur != nil {
		cur = cur.Next
		lastK--
	}
	if lastK > 0 {
		return h
	} else if lastK == 0 {
		return h.Next
	} else {
		cur = h
		lastK++
		for lastK < 0 {
			cur = cur.Next
			lastK++
		}
		cur.Next = cur.Next.Next
		return h
	}
}

func RemoveLastKthNodeDouble(h *DoubleNode, lastK int) *DoubleNode {
	if h == nil {
		return nil
	}
	cur := h
	for cur != nil {
		cur = cur.Next
		lastK--
	}
	if lastK == 0 {
		h = h.Next
	}
	if lastK < 0 {
		cur = h
		lastK++
		for lastK < 0 {
			cur = cur.Next
			lastK++
		}
		cur.Next = cur.Next.Next
		if cur.Next != nil {
			cur.Next.Last = cur
		}
	}
	return h
}
