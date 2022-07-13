package list

import "math"

func RemoveByMid(h *Node) *Node {
	if h == nil {
		return nil
	}
	if h.Next == nil {
		return h
	}
	if h.Next.Next == nil {
		return h.Next
	}
	cur := h.Next.Next
	pre := h
	for cur.Next != nil && cur.Next.Next != nil {
		cur = cur.Next.Next
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return h
}

func RemoveByRatio(h *Node, a, b int) *Node {
	if h == nil || h.Next == nil {
		return nil
	}
	cur := h
	n := 0
	for cur != nil {
		cur = cur.Next
		n++
	}
	k := math.Ceil((float64(n) * float64(a)) / float64(b))
	kn := int(k)
	if kn == 1 {
		h = h.Next
	}
	if kn > 1 {
		kn--
		pre := h
		for kn > 1 {
			pre = pre.Next
			kn--
		}
		pre.Next = pre.Next.Next
	}
	return h
}
