package list

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

func CloneRandom(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		clone := &RandomListNode{
			Label: cur.Label,
			Next:  cur.Next,
		}
		cur.Next = clone
		cur = cur.Next.Next
	}
	newHead := head.Next
	cur = head
	clone := cur.Next
	for cur != nil {
		if cur.Random != nil {
			clone.Random = cur.Random.Next
		}
		if cur.Next != nil {
			cur = cur.Next.Next
		}
		if clone.Next != nil {
			clone = clone.Next.Next
		}
	}

	cur = head
	clone = cur.Next
	for cur != nil {
		if cur.Next != nil {
			cur.Next = cur.Next.Next
		}
		if clone.Next != nil {
			clone.Next = clone.Next.Next
		}
	}
	return newHead
}
