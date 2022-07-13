package list

func ListPartition(h *Node, pivot int) *Node {
	if h == nil {
		return nil
	}
	var sH *Node
	var sT *Node
	var eH *Node
	var eT *Node
	var bH *Node
	var bT *Node
	cur := h
	for cur != nil {
		if cur.Val < pivot {
			if sH == nil {
				sH = cur
				sT = cur
			} else {
				sT.Next = cur
				sT = cur
			}
		} else if cur.Val == pivot {
			if eH == nil {
				eH = cur
				eT = cur
			} else {
				eT.Next = cur
				eT = cur
			}
		} else {
			if bH == nil {
				bH = cur
				bT = cur
			} else {
				bT.Next = cur
				bT = cur
			}
		}
		cur = cur.Next
	}
	if sT != nil {
		sT.Next = eH
		if eH == nil {
			eT = sT
		}
	}
	if eT != nil {
		eT.Next = bH
	}
	if sH != nil {
		return sH
	} else if eH != nil {
		return eH
	}
	return bH
}
