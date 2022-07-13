package list

func DeleteDuplicate(h *Node) *Node {
	if h == nil {
		return nil
	}
	pre := &Node{
		Val:  0,
		Next: h,
	}
	cur := pre

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			tmp := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == tmp {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}

	}
	return pre.Next
}
