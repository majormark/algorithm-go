package list

func isPalindrome(h *Node) bool {
	if h == nil {
		return false
	}
	if h.Next == nil {
		return true
	}
	slow := h
	fast := h
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	node1 := slow.Next
	pre := slow
	for node1 != nil {
		next := node1.Next
		node1.Next = pre
		pre = node1
		node1 = next
	}
	slow.Next = nil
	last := pre
	first := h
	res := true
	for first != nil {
		if first.Val != last.Val {
			res = false
			break
		}
		first = first.Next
		last = last.Next
	}

	last = pre
	pre = nil
	for last != nil {
		next := last.Next
		last.Next = pre
		pre = last
		last = next
	}
	return res
}
