package list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Convert1(h *TreeNode) *TreeNode {
	if h == nil {
		return nil
	}
	queue := []*TreeNode{}
	inOrder(queue, h)
	h = queue[0]
	queue = queue[1:]
	pre := h
	var cur *TreeNode
	for len(queue) > 0 {
		cur = queue[0]
		queue = queue[1:]
		pre.Right = cur
		cur.Left = pre
		pre = cur
	}
	pre.Right = nil
	return h
}

func inOrder(queue []*TreeNode, node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(queue, node.Left)
	queue = append(queue, node)
	inOrder(queue, node.Right)
}
