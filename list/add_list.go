package list

func AddList(h1 *Node, h2 *Node) *Node {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	h1Reverse := reverseList(h1)
	h2Reverse := reverseList(h2)
	var newHead *Node
	var ca int
	var n1 int
	var n2 int
	for h1Reverse != nil || h2Reverse != nil {
		if h1Reverse == nil {
			n1 = 0
		} else {
			n1 = h1Reverse.Val
		}
		if h2Reverse == nil {
			n2 = 0
		} else {
			n2 = h2Reverse.Val
		}
		val := n1 + n2 + ca
		node := &Node{}
		node.Val = val % 10
		ca = node.Val / 10
		node.Next = newHead
		newHead = node
		if h1Reverse != nil {
			h1Reverse = h1Reverse.Next
		}
		if h2Reverse != nil {
			h2Reverse = h2Reverse.Next
		}
	}

	if ca == 1 {
		node := &Node{
			Val:  1,
			Next: newHead,
		}
		newHead = node
	}
	reverseList(h1Reverse)
	reverseList(h2Reverse)
	return newHead
}

func reverseList(h *Node) *Node {
	if h == nil {
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
