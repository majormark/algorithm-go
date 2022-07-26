package list

func SelectionSort(h *Node) *Node {
	if h == nil || h.Next == nil {
		return h
	}
	cur := h
	var newHead *Node
	var newTail *Node
	var min *Node
	for cur != nil {
		min = cur
		innerCur := cur.Next
		pre := cur
		preMin := cur
		for innerCur != nil {
			if innerCur.Val < min.Val {
				min = innerCur
				preMin = pre
			}
			pre = innerCur
			innerCur = innerCur.Next
		}
		if min != cur {
			preMin.Next = min.Next
		} else {
			cur = cur.Next
		}

		if newTail == nil {
			newTail = min
		} else {
			newTail.Next = min
		}
		if newHead == nil {
			newHead = min
		}
	}
	return newHead
}
