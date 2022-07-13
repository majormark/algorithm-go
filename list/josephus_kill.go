package list

func JosephusKill1(h *Node, m int) *Node {
	if h == nil || m < 1 {
		return h
	}
	last := h
	for last.Next != h {
		last = last.Next
	}
	count := 0
	for last != h {
		count++
		if count == m {
			last.Next = h.Next
			count = 0
		} else {
			last = last.Next
		}
		h = last.Next
	}
	return last
}
