package intersection

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

func GetLoopNode(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	slow := head.Next
	fast := head.Next.Next

	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func NoLoop(h1 *Node, h2 *Node) *Node {
	if h1 == nil || h2 == nil {
		return nil
	}
	end1, len1 := findEndNodeAndLen(h1)
	end2, len2 := findEndNodeAndLen(h2)

	if end1 != end2 {
		return nil
	}
	tmp1 := h1
	tmp2 := h2
	if len1 > len2 {
		n := len1 - len2
		for i := 0; i < n; i++ {
			tmp1 = tmp1.Next
		}

	} else {
		n := len2 - len1
		for i := 0; i < n; i++ {
			tmp2 = tmp2.Next
		}
	}

	for tmp1 != tmp2 {
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
	}
	return tmp1
}

func findEndNodeAndLen(h *Node) (*Node, int) {
	tmp := h
	length := 0
	var end *Node
	for tmp != nil {
		if tmp.Next == nil {
			end = tmp
		}
		tmp = tmp.Next
		length++
	}
	return end, length
}

func BothLoop(h1 *Node, loop1 *Node, h2 *Node, loop2 *Node) *Node {
	if h1 == nil || loop1 == nil || h2 == nil || loop2 == nil {
		return nil
	}
	if loop1 == loop2 {
		tmp1 := h1
		tmp2 := h2
		n := 0
		for tmp1 != loop1 {
			tmp1 = tmp1.Next
			n++
		}
		for tmp2 != loop2 {
			tmp2 = tmp2.Next
			n--
		}
		cur1 := h1
		cur2 := h2
		if n < 0 {
			cur1 = h2
			n = -n
		}
		if cur1 == h2 {
			cur2 = h1
		}
		for i := 0; i < n; i++ {
			cur1 = cur1.Next
		}
		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1
	} else {
		tmp1 := h1
		for tmp1 != loop1 {
			if tmp1 == loop2 {
				return loop1
			}
			tmp1 = tmp1.Next
		}
		return nil
	}
}

func FindIntersection(h1 *Node, h2 *Node) *Node {
	if h1 == nil || h2 == nil {
		return nil
	}
	loop1 := GetLoopNode(h1)
	loop2 := GetLoopNode(h2)
	if loop1 == nil && loop2 == nil {
		return NoLoop(h1, h2)
	}
	if loop1 != nil && loop2 != nil {
		return BothLoop(h1, loop1, h2, loop2)
	}
	return nil
}
